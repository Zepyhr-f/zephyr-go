# Zephyr Admin (Go)

> 基于 **go-zero** 微服务框架构建的中后台管理系统后端，采用 **BFF 网关 + 内部 gRPC 微服务** 架构，收敛于单一 Monorepo 仓库。
>
> 当前已落地：统一认证（双 Token）、RBAC 身份体系（用户 / 角色 / 菜单 / 部门 / 岗位）、后台运维扩展能力、服务日志监控，以及对 `zephyr-ai` 的反向代理与网关签名穿透。

---

## 目录

- [项目简介](#项目简介)
- [技术栈](#技术栈)
- [系统架构](#系统架构)
- [仓库结构](#仓库结构)
- [核心服务](#核心服务)
- [功能特性](#功能特性)
- [快速开始](#快速开始)
- [配置说明](#配置说明)
- [API 接口概览](#api-接口概览)
- [认证与安全机制](#认证与安全机制)
- [代码生成（goctl）](#代码生成goctl)
- [数据库](#数据库)
- [项目文档](#项目文档)
- [开发规范](#开发规范)

---

## 项目简介

`zephyr-go` 是 Zephyr Admin 平台的 Go 后端重构版本，定位为可演进的中后台底座。整体遵循 **契约优先（Schema-First）** 的开发范式：先编写 `.api` / `.proto` 契约，再用 `goctl` 生成骨架代码，最后在 `logic` 层填充业务编排。

系统划分为以下域：

| 域 | 服务 | 状态 | 说明 |
| --- | --- | --- | --- |
| 网关层 | `app/gateway` | ✅ 已实现 | BFF 聚合网关，统一对外 HTTP 入口 |
| 统一认证域 | `app/auth` | ✅ 已实现 | 账密登录、双 Token 签发、安全登出（gRPC） |
| 身份体系域 | `app/identity` | ✅ 已实现 | 用户 / 角色 / 菜单 / 部门 / 岗位 RBAC（gRPC） |
| AI 核心能力域 | `zephyr-ai`（外部） | 🔗 已接入 | 通过网关反向代理 + HMAC 签名穿透 |
| 多租户 / RAG / Agent / Workflow | — | 🗺 规划中 | 详见 [doc/架构设计](doc/架构设计) |

---

## 技术栈

| 分类 | 选型 |
| --- | --- |
| 语言 | Go 1.26 |
| 微服务框架 | [go-zero](https://github.com/zeromicro/go-zero) v1.10 |
| 代码生成 | goctl 1.10 |
| RPC | gRPC + Protobuf |
| HTTP | go-zero rest |
| ORM | GORM + PostgreSQL（pgvector） |
| 缓存 / 会话 | Redis（go-zero stores/redis） |
| 认证 | JWT（HS256，golang-jwt/jwt/v4）+ bcrypt |
| 可观测 | OpenTelemetry / Prometheus（go-zero 内置） |
| 容器化 | Docker + docker-compose |

---

## 系统架构

```text
全局请求入口 (Nginx / Load Balancer)
      ↓  HTTP
API BFF 网关 (app/gateway)  :8888
  · 统一认证(JWT) + 黑名单
  · URL-RBAC 粗粒度鉴权
  · 统一响应封装 R<T> + 全局异常
  · AI 反向代理 + 网关 HMAC 签名穿透
  · 服务日志监控（只读 / SSE tail）
      ↓  gRPC
后端核心微服务集群
  ├── 统一认证域  (app/auth)       :8080  · 双 Token 签发、登出黑名单
  └── 身份体系域  (app/identity)   :8081  · RBAC 数据 / 数据权限
      ↓
数据访问层 (GORM → PostgreSQL / Redis)
      ↓
AI 上游 (zephyr-ai)  · 经网关反向代理，注入 X-Zephyr-* 可信 Header + 网关签名
```

**部署拓扑（docker-compose）：**

```text
infra_default (external network)
  ├── pgsql-pgvector  :5432   PostgreSQL + pgvector
  ├── redis           :6379   会话 / RefreshToken / 黑名单 / 鉴权缓存
  ├── zephyr-identity:8081   依赖 PostgreSQL
  ├── zephyr-auth    :8080   依赖 identity、Redis
  └── zephyr-gateway  :8888   依赖 auth、identity；反代 zephyr-ai
```

> Redis 与 PostgreSQL 运行于外部网络 `infra_default`，需由独立的基建 compose 提供（生产密钥通过 `docker-compose.override.yml` 覆盖，已 gitignore）。

---

## 仓库结构

```text
zephyr-go/
├── go.mod / go.sum            # Monorepo 统一依赖管理
├── docker-compose.yml         # 三服务编排（identity / auth / gateway）
├── app/                       # 应用层：网关与各微服务
│   ├── gateway/               # ① BFF 网关（HTTP 入口）
│   │   ├── gateway.go         #   main 启动入口
│   │   ├── gateway.api        #   HTTP API 主入口（import 聚合 desc/*.api）
│   │   ├── desc/              #   按业务域拆分的 API 契约（auth/identity/user/role/menu/dept/post）
│   │   ├── etc/               #   gateway-api.yaml / gateway-api-docker.yaml
│   │   └── internal/          #   config / handler / logic / svc / types
│   ├── auth/                  # ② 统一认证域（gRPC）
│   │   ├── auth.go            #   main 启动入口
│   │   ├── zephyr_auth.proto  #   gRPC 契约
│   │   ├── pb/                #   goctl 生成桩代码
│   │   ├── etc/               #   auth.yaml / auth-docker.yaml
│   │   ├── gen_pwd.go         #   bcrypt 口令生成工具
│   │   └── internal/          #   config / logic / server / svc
│   └── identity/              # ③ 身份体系域（gRPC）
│       ├── identity.go        #   main 启动入口
│       ├── identity.proto     #   gRPC 契约（User/Role/Menu/Dept/Post）
│       ├── pb/                #   goctl 生成桩代码
│       ├── etc/               #   identity.yaml / identity-docker.yaml
│       └── internal/          #   config / logic / server / svc / repository(model)
├── pkg/                       # 公共库（zephyr-framework 的 Go 平替）
│   ├── core/                  #   统一响应 R<T>、业务错误码 xerr
│   ├── orm/                   #   GORM 基座（BaseModel：审计字段 + 软删除）
│   ├── gatewaysign/           #   网关 HMAC 签名协议（X-Zephyr-* Header / Nonce 存储）
│   └── logmask/               #   日志脱敏
└── doc/                       # 架构 / 规范 / 数据库设计文档
```

---

## 核心服务

### app/gateway — BFF 聚合网关

- **职责**：前端唯一入口。统一认证（JWT）、URL-RBAC 粗粒度鉴权、协议转换（HTTP ↔ gRPC）、统一响应封装、全局异常处理。
- **入口**：`gateway.go`，HTTP 监听 `:8888`。
- **契约分片**：`gateway.api` 通过 `import "desc/*.api"` 聚合各业务域接口，避免单文件膨胀。
- **AI 反向代理**：`/api/v1/ai/*` 反代至 `zephyr-ai`，剥离 `Authorization`/`Cookie`，注入 `X-Zephyr-*` 可信 Header 与网关 HMAC 签名（`pkg/gatewaysign`）。
- **服务日志监控**：`/api/v1/sysadmin/logs/*` 提供只读日志浏览 / 搜索 / 下载 / SSE 实时 tail（SSE 路由单独注册 `WithTimeout(0)` 以支持流式）。

### app/auth — 统一认证域（gRPC）

- **职责**：账号密码登录校验、双 Token 签发、安全登出。
- **入口**：`auth.go`，gRPC 监听 `:8080`（Dev/Test 模式开启 reflection）。
- **登录链路**：`LoginVerify` → 调 `identity.GetUserAuthInfo` 取密码哈希 → `bcrypt` 校验 → 签发 **Access Token（JWT HS256）** + **Refresh Token（不透明 UUID，存 Redis，TTL 7 天）**。
- **登出**：将 Token `jti` 写入 Redis 黑名单，TTL 为 Access Token 剩余有效期。

### app/identity — 身份体系域（gRPC）

- **职责**：RBAC 全量数据 CRUD 与权限分配。
- **入口**：`identity.go`，gRPC 监听 `:8081`。
- **能力**：用户、角色、菜单（树）、部门（树）、岗位的增删改查；角色-菜单分配（`RoleAssignMenus`）、数据权限范围（`RoleDataScope`）；当前用户信息 + 角色 + 按钮级 perms（`GetUserInfo`）。
- **持久化**：GORM + PostgreSQL，模型定义于 `internal/repository/model`（`sys_dept` / `sys_menu` / `sys_role` …）。

### pkg — 公共库

| 包 | 说明 |
| --- | --- |
| `pkg/core/response` | 统一响应体 `{code, msg, data}`（`R<T>`），`Success` / `Error` 封装 |
| `pkg/core/xerr` | 统一业务错误码定义 |
| `pkg/orm` | `BaseModel` 基座：`ID` / `CreatedAt` / `UpdatedAt` / `CreatedBy` / `UpdatedBy` / `DeletedAt`（软删除） |
| `pkg/gatewaysign` | 网关签名协议：`X-Zephyr-Gateway-Sign/Ts/Nonce`、`X-Zephyr-User-Id/Username/Tenant-Id/Roles`；HMAC-SHA256 签名 / 常量时间校验；Nonce 存储（Redis / 内存双实现） |
| `pkg/logmask` | 敏感字段日志脱敏 |

---

## 功能特性

- **双 Token 无感刷新**：Access Token（JWT，短）+ Refresh Token（HttpOnly Cookie，长，可吊销）。
- **登出黑名单**：基于 Redis `jti` 黑名单，支持主动失效。
- **RBAC 权限体系**：用户 ↔ 角色 ↔ 菜单（按钮 perms），支持角色-菜单分配与数据范围（DataScope）。
- **网关签名穿透**：内部服务可信 Header + HMAC 防伪造，AI 上游验签后即可信任身份。
- **后台运维扩展**：登录日志 / 操作日志 / 在线用户（踢出）、服务与缓存监控、定时任务管理、字典 / 参数 / 文件 / 通知、代码生成预览、SQL 执行。
- **服务日志监控**：跨服务日志只读浏览与实时 tail。
- **统一响应与异常**：网关层收敛 gRPC 错误为前端可读的 `R<T>` 结构。

---

## 快速开始

### 前置依赖

- Go ≥ 1.26
- [goctl](https://github.com/zeromicro/go-zero) v1.10（代码生成）
- PostgreSQL（建议带 pgvector 扩展）
- Redis
- Docker & docker-compose（容器化部署）

### 1. 准备数据库

按 [doc/数据库设计](doc/数据库设计) 的执行顺序初始化：

1. `4.2.1-infrastructure.sql` — 基础数据库环境
2. `4.1.x-*.sql` — 建表脚本（rbac_core / security_audit / platform_config / sys_job / infrastructure_module）
3. `4.2.2+` 系列 — 初始化权限与默认数据

生成初始管理员口令哈希（用于写入 `sys_user`）：

```bash
cd app/auth && go run gen_pwd.go   # 输出 admin123 的 bcrypt 哈希
```

### 2. 本地运行（三服务 + 依赖）

> 各服务默认读取 `etc/*.yaml`（本地直连 `127.0.0.1`）。请先按需修改 `app/*/etc/*.yaml` 中的数据库 DSN、Redis、JWT Secret。

```bash
# 终端 1：身份服务
cd app/identity && go run identity.go -f etc/identity.yaml

# 终端 2：认证服务
cd app/auth     && go run auth.go     -f etc/auth.yaml

# 终端 3：网关
cd app/gateway  && go run gateway.go  -f etc/gateway-api.yaml
```

启动后：

| 服务 | 地址 |
| --- | --- |
| Gateway HTTP | http://localhost:8888 |
| Auth gRPC | 127.0.0.1:8080 |
| Identity gRPC | 127.0.0.1:8081 |

### 3. Docker 部署

容器化配置读取 `etc/*-docker.yaml`（服务间以容器名寻址，日志落盘 `/app/logs/<svc>`）。需先在 `infra_default` 网络中提供 `pgsql-pgvector` 与 `redis`，并准备 `docker-compose.override.yml` 填入真实密钥：

```bash
docker compose up -d --build
```

---

## 配置说明

每个服务提供两套配置：本地 `etc/*.yaml` 与容器 `etc/*-docker.yaml`。关键字段如下。

**Gateway（`gateway-api.yaml`）**

| 字段 | 说明 |
| --- | --- |
| `Host` / `Port` | HTTP 监听地址（默认 `0.0.0.0:8888`） |
| `AuthRpc.Target` / `IdentityRpc.Target` | 下游 gRPC 地址 |
| `Auth.AccessSecret` / `AccessExpire` | JWT 校验密钥与有效期（须与 auth 一致） |
| `BizRedis` | Redis 连接（鉴权缓存等） |
| `Gateway.SignSecret` / `SignTTLSeconds` | 网关 HMAC 签名密钥与防重放窗口 |
| `Gateway.AiUpstream` / `AiPathPrefix` / `AiUpstreamStrip` | AI 反代上游与路径改写 |

**Auth（`auth.yaml`）**

| 字段 | 说明 |
| --- | --- |
| `ListenOn` | gRPC 监听地址（默认 `:8080`） |
| `IdentityRpc.Target` | identity 服务地址 |
| `BizRedis` | Refresh Token / 黑名单存储 |
| `JwtAuth.AccessSecret` / `AccessExpire` | JWT 签发密钥与有效期 |

**Identity（`identity.yaml`）**

| 字段 | 说明 |
| --- | --- |
| `ListenOn` | gRPC 监听地址（默认 `:8081`） |
| `Postgres.DSN` | PostgreSQL 连接串 |

> ⚠️ 提交的 `*-docker.yaml` 仅为示例，真实密钥请通过 `docker-compose.override.yml` 或 `app/*/etc/*-docker.local.yaml` 覆盖（均已 gitignore）。

---

## API 接口概览

网关对外暴露的 HTTP 接口（详见 [app/gateway/desc](app/gateway/desc)）：

| 模块 | 前缀 | 鉴权 | 主要接口 |
| --- | --- | --- | --- |
| 认证 | `/api/v1/auth` | 白名单 | `POST /login`、`POST /refresh`、`POST /logout` |
| 当前用户 | `/api/v1/auth` | JWT | `GET /info`（用户 + roles + perms） |
| 用户管理 | `/api/v1/system/user` | JWT | list / submit / updateStatus / resetPassword / remove |
| 角色管理 | `/api/v1/system/role` | JWT | list / submit / updateStatus / remove / detail / menuTree / assignMenus / dataScope |
| 菜单管理 | `/api/v1/system/menu` | JWT | tree / list / detail / save / update / remove / status |
| 部门管理 | `/api/v1/system/dept` | JWT | tree / save / update / remove |
| 岗位管理 | `/api/v1/system/post` | JWT | list / save / update / status / remove |
| 安全审计 | `/api/v1/security` | JWT | 登录日志 / 操作日志 / 在线用户（踢出） |
| 监控 | `/api/v1/monitor` | JWT | server / cache / datasource 指标、定时任务管理 |
| 基础设施 | `/api/v1/infrastructure` | JWT | 字典 / 参数 / 文件 / 通知 |
| 开发工具 | `/api/v1/devtools` | JWT | 代码生成 / API 文档 / SQL 执行 |
| AI 代理 | `/api/v1/ai/*` | `/health` 公开，其余 JWT | 反代 zephyr-ai（注入签名） |
| 日志监控 | `/api/v1/sysadmin/logs` | JWT | 服务 / 文件 / 搜索 / 下载 / SSE tail |

**统一响应格式：**

```json
{ "code": 200, "msg": "success", "data": { } }
```

---

## 认证与安全机制

```text
前端  ──Authorization: Bearer {accessToken}──▶  Gateway
                                                  ├─ JWT 校验 + 黑名单
                                                  ├─ URL-RBAC 粗粒度鉴权
                                                  ├─ 注入 X-Zephyr-* 可信 Header
                                                  └─ 注入 X-Zephyr-Gateway-Sign (HMAC)
                                                            │
                                                            ▼
                                                  后端服务 / zephyr-ai
                                                  验签 → 信任 Header → 细粒度 perms + DataScope
```

- **双 Token**：Access = JWT（短，HS256，`userCode` / `tenantCode` / `jti`）；Refresh = 不透明 UUID，存 Redis（TTL 7 天），网关剥离后置入 HttpOnly Cookie，防 XSS 窃取。
- **登出黑名单**：`blacklist:{jti}` 写入 Redis，TTL 为 Access 剩余有效期。
- **网关签名**：`pkg/gatewaysign` 以 `METHOD \n path \n ts \n nonce \n hex(sha256(body))` 为 canonical message，HMAC-SHA256 签名，常量时间校验；配合时间戳 + Nonce 防重放。
- **职责边界**：网关负责认证与粗粒度鉴权；服务侧验签后信任 Header，执行按钮级 perms 与数据权限 DataScope。

> 完整链路口径以 [doc/架构设计/02-安全与权限](doc/架构设计/02-安全与权限) 为单一事实来源（SSOT）。

---

## 代码生成（goctl）

项目以契约优先驱动，修改 `.api` / `.proto` 后重新生成骨架：

```bash
# HTTP 网关：在 app/gateway 下执行
goctl api go -api gateway.api -dir .

# gRPC 服务（以 identity 为例）：在 app/identity 下执行
goctl rpc protoc identity.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

# gRPC 服务（以 auth 为例）：在 app/auth 下执行
goctl rpc protoc zephyr_auth.proto --go_out=./pb/pb --go-grpc_out=./pb/pb --zrpc_out=.
```

生成后只需在 `internal/logic` 填充业务编排，鉴权 / 路由 / 参数解析 / 统一响应由框架与底座代理。

---

## 数据库

- **引擎**：PostgreSQL（生产使用带 pgvector 扩展的实例，为后续 RAG 向量检索预留）。
- **ORM**：GORM，基座 `pkg/orm.BaseModel`（审计字段 + 软删除）。
- **脚本**：集中于 [doc/数据库设计](doc/数据库设计)，按 `4.2.1 → 4.1.x → 4.2.2+` 顺序执行。

| 脚本 | 关键词 | 说明 |
| --- | --- | --- |
| `4.1.1-rbac_core.sql` | user / role / menu | RBAC 核心表 |
| `4.1.2-security_audit.sql` | audit / log | 安全审计表 |
| `4.1.3-platform_config.sql` | dict / config | 平台配置 / 字典 |
| `4.1.3-sys_job.sql` | job / scheduler | 定时任务 |
| `4.1.4-infrastructure_module.sql` | infra / module | 基础设施模块 |

---

## 项目文档

详细设计文档位于 [doc/](doc)，建议按以下顺序阅读：

- **架构设计**
  - [01-系统架构](doc/架构设计/01-系统架构) — 工程结构 / Gateway / Identity / Auth 设计方案
  - [02-安全与权限](doc/架构设计/02-安全与权限) — 登录鉴权全链路、安全基线 / 数据权限 / 审计
  - [03-数据与存储](doc/架构设计/03-数据与存储) — 数据模型、迁移初始化、缓存策略
- **开发规范** — [建表规范](doc/开发规范/00-建表规范.md)、[AI 协同开发规范](doc/开发规范/05-AI协同开发规范.md)
- **数据库设计** — [建表脚本](doc/数据库设计/01-建表脚本)、[初始化脚本](doc/数据库设计/02-初始化脚本)

---

## 开发规范

- **契约优先**：接口变更先改 `.api` / `.proto`，再 `goctl` 重新生成。
- **分层约束**：`logic` 层只做用例编排，禁止直接写 SQL；GORM 操作下沉至 `repository`，外部调用收拢至 `adapter`，经 `svc` 依赖注入。
- **统一响应**：所有 HTTP 出口使用 `pkg/core/response` 的 `R<T>` 结构，业务错误走 `xerr`。
- **配置安全**：密钥不进版本库，通过 override 文件注入。
- 提交规范见 [doc/开发规范/03-Git工作流与提交规范.md](doc/开发规范/03-Git工作流与提交规范.md)。
