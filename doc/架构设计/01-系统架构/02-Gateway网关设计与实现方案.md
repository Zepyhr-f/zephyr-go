# BFF 网关 (Gateway) 设计与实现方案

## 1. 定位与职责
作为系统微服务架构的唯一外部流量入口，Gateway（BFF层，Backend For Frontend）不再承担核心业务的计算和持久化，而是专注于以下通用非业务功能和接口编排：
- **统一接入**：接收所有前端（Web/App）请求，对外暴露 HTTP RESTful 接口。
- **协议转换**：将外部的 HTTP 请求按需聚合或转换为内部微服务的 gRPC 请求。
- **安全与鉴权**：负责 JWT Token 的签发验证、URL 级别的权限拦截（RBAC 防线）。
- **统一响应与异常处理**：收敛后端微服务返回的 gRPC 错误，统一转换为前端可读的统一 JSON 格式（`R<T>`）与业务状态码。
- **限流与熔断**：基于 go-zero 的原生中间件机制，保护后端核心服务不受突发流量冲垮。

---

## 2. 目录架构
网关代码集中于 `app/gateway` 目录下，采用 API 驱动开发（Schema-First）：

```text
app/gateway/
├── etc/
│   └── gateway.yaml           # 网关服务配置（端口、ETCD发现、Redis缓存、JWT密钥等）
├── gateway.api                # API 主配置入口，利用 import 聚合各业务域
├── desc/                      # 按业务域拆分的独立 API 定义文件
│   ├── auth.api               # 登录/登出/Token刷新相关的公共接口 (内部调用 identity 微服务)
│   ├── identity.api           # 用户管理/角色管理/菜单管理 (内部调用 identity 微服务)
│   └── agent.api              # AI Agent 调度接口 (内部调用 agent 微服务)
├── internal/
│   ├── config/                # [自动生成] 对应 gateway.yaml 的强类型结构体
│   ├── handler/               # [自动生成] HTTP 请求路由与参数解析绑定
│   ├── logic/                 # [自动生成] API 编排逻辑层，负责聚合/调用下游多个微服务 RPC 
│   ├── middleware/            # [手动实现] 自定义全局及路由级中间件（URL鉴权、操作日志审计等）
│   ├── svc/                   # 依赖注入容器（RPC Clients, Redis连接, Middleware 等初始化）
│   └── types/                 # [自动生成] HTTP 请求与响应的 DTO 强类型结构
└── gateway.go                 # 网关 main 函数启动入口
```

---

## 3. 核心机制实现方案

### 3.1 API 契约分片管理
为了解决随着业务增加导致单个 `.api` 文件动辄数千行难以维护的问题，我们规范化利用 `go-zero` 的 import 特性。

主入口文件 `gateway.api` 示例：
```go
syntax = "v1"

info(
    title: "Zephyr Admin API Gateway"
    desc: "统一网关总入口"
    author: "Zephyr"
    version: "1.0.0"
)

// 按业务域导入拆分后的子模块
import "desc/auth.api"
import "desc/identity.api"
import "desc/tenant.api"
import "desc/agent.api"
import "desc/workflow.api"
```
**规范**：每次新增核心模块或微服务时，必须在 `desc/` 目录下新建同名 `.api` 文件，并在主入口注册，保持网关代码生成的模块化。

### 3.2 JWT 鉴权与双 Token 机制
网关作为安全的第一道防线，需要解析请求携带的凭证，并隐式传递给内部微服务，避免下游微服务重复鉴权。

1. **网关拦截与解析**：
   在对应的 `.api` 文件分组中加上 `@server(jwt: Auth)`，网关会自动拦截 HTTP Header 中的 `Authorization: Bearer <Token>`。验证通过后，将 Token 携带的 payload（如 `userCode`, `tenantCode`）塞入 `context.Context` 中。

2. **双 Token 无感刷新与跨域安全控制**：
   为兼容前端体系并提供最高级别安全性，Identity 服务会签发 `AccessToken` 和 `RefreshToken`，但 **网关负责剥离 `RefreshToken` 并将其注入到 `HttpOnly Secure Cookie` 中响应给前端**。当前端触发 `/auth/refresh` 接口时，网关读取 Cookie 中的刷新令牌传给下游重置，彻底防御 XSS 窃取。

3. **RPC 级联透传**：
   在 `logic` 层发起 RPC 调用前，从上下文中取出 `userCode` 等标识，通过 gRPC 的 `metadata.AppendToOutgoingContext` 附加到传输头中向下游隐式透传。

### 3.3 统一响应体格式化封装
传统前端要求标准的响应体 `{"code": 200, "msg": "ok", "data": {...}}`，而 gRPC 内部调用的返回是纯数据实体或特定的 `gRPC Error`。网关需要在此交汇点实现统一拦截包装。

- **通用成功包装**：
  在 `pkg/core/response` 包中提供通用的 `response.Success(w, resp)`，在 `handler` 层面替换 `go-zero` 生成的默认的 `httpx.OkJson` 返回方式。
- **全局异常拦截**：
  在 `gateway.go` 的启动流程中，通过 `httpx.SetErrorHandler` 注册全局错误处理器。当捕捉到 `gRPC status.Error` 时，解析出我们在内部定义的自定义错误码（xerr），转化为统一的 HTTP 业务错误结构响应给前端。

### 3.4 基于路由权限的动态鉴权 (RBAC 与 Redis 联动)
合法登录（拥有有效 JWT）只是基础校验。对于绝大多数后台管理接口，还需判断当前用户是否有权限访问该接口（例如：普通员工禁止访问 `POST /api/v1/role/create`）。

- **实现方式**：自定义路由中间件 `AuthorityMiddleware`。在 `.api` 文件中通过 `@server(middleware: AuthorityMiddleware)` 挂载到受保护的路由组。
- **处理流**：
  1. 获取当前请求的 `Path` 和 `Method`（作为资源标识符）。
  2. 从 `Context` 中提取当前用户的身份标识（如 `tenantCode` 与 `userCode`）。
  3. 极速读取 Redis 中的高频鉴权缓存 `zephyr:auth:user_roles:<tenant_code>:<user_code>` (该 Set 由 Identity 在登录/授权时维护)。
  4. 如果命中角色白名单或权限路由网格，放行请求 `next(w, r)`；如果越权未匹配，直接在网关层抛弃，返回 `403 Forbidden`。

---

## 4. 开发工作流 (SOP)
基于上述规范，在网关层开发一个新的接入接口（例如：获取用户列表）的标准流程如下：

1. **定义 API 契约**：在 `desc/identity.api` 中编写 `/user/list` 请求和响应体结构。
2. **生成骨架代码**：在 `app/gateway` 目录下执行编译指令：
   ```bash
   goctl api go -api gateway.api -dir .
   ```
3. **注入 RPC 客户端**：在 `internal/svc/servicecontext.go` 中注册 `IdentityRpc` 客户端（由内部微服务提供）。
4. **填充业务编排**：打开 `internal/logic/identity/getuserlistlogic.go`，将 HTTP 参数转化为 gRPC 请求，发起 RPC 调用，然后原样或者转换格式后返回数据。
5. **无需操心周边逻辑**：鉴权、拦截、JSON 统一结构包装均已由框架底座和中间件代理，研发只需专注步骤 4 中的 RPC 编排即可。
