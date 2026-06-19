# Identity 身份服务设计与实现方案

## 1. 领域边界与职责
Identity 服务（`app/identity`）是基础认证鉴权与组织架构的提供者，属于微服务体系的核心基座。
根据数据库 PgSQL 的设计及原有 Java `zephyr-server` 的业务逻辑，该微服务负责管理以下领域对象：

- **用户子域 (User)**: `zephyr_sys_user`, `zephyr_sys_user_role`, `zephyr_sys_user_post`
- **角色与权限子域 (Role & Menu)**: `zephyr_sys_role`, `zephyr_sys_menu`, `zephyr_sys_role_menu`, `zephyr_sys_role_dept`
- **组织架构子域 (Dept & Post)**: `zephyr_sys_dept`, `zephyr_sys_post`

*(注：`zephyr_sys_tenant` 表划归 `app/tenant` 租户微服务，但在本域中各业务表均保留 `tenant_code` 字段以配合数据隔离逻辑)*

---

## 2. Redis 缓存状态机设计 (兼容 Java 逻辑)
结合原系统 `AuthController.java`，身份域的认证并不仅仅查库，还需要重度依赖 Redis 维护会话：

- **Refresh Token 维护**：存入 `zephyr:auth:refresh:<jti>`，实现 7 天内无感刷新 AccessToken 的能力。
- **用户全局信息缓存**：`zephyr:auth:user_info:<tenant_code>:<user_code>` (过期时间2小时)，内部缓存着该用户的完整角色与权限列表，加速内部微服务读取，减轻数据库压力。
- **角色白名单缓存**：`zephyr:auth:user_roles:<tenant_code>:<user_code>` (Set结构)，专供网关层的 RBAC 拦截器极速鉴定请求权限。

---

## 3. GORM 模型设计与 PgSQL 特色适配

原始的 Java/MyBatis 架构迁移至 Go/GORM 时，针对 `04-数据库设计` 中的 PgSQL 特性，我们需要做针对性适配：

### 3.1 自定义软删除 (Soft Delete)
数据库建表脚本中未采用传统的时间戳 `deleted_at` 字段作为软删除，而是使用了：
```sql
"if_deleted" SMALLINT NOT NULL DEFAULT 0
CONSTRAINT uk_code_del UNIQUE ("code", "if_deleted", "tenant_code")
```
> [!IMPORTANT]
> **设计决议**：不能使用 `gorm.DeletedAt` 魔法字段。
> - 在 PO 模型中手动定义 `IfDeleted int16`。
> - 开发自定义 GORM Plugin (或者在 Dao 层封装全局 Scope) 注入 `if_deleted = 0` 的查询条件。
> - 执行删除操作时，必须封装为 `db.Model(&po).Update("if_deleted", 1)`。

### 3.2 业务主键外键关联 (Code-based Association)
主物理主键是雪花算法 `id`，但各表之间的关联表（如 `sys_user_role`）使用的是 `user_code` 和 `role_code`。
> [!TIP]
> **设计决议**：
> GORM 需要在关联关系中使用 `references` 和 `joinReferences` 标签强制指定业务外键：
```go
type SysUser struct {
    Id         int64  `gorm:"primaryKey"`
    Code       string `gorm:"uniqueIndex"`
    Roles      []SysRole `gorm:"many2many:zephyr_sys_user_role;foreignKey:Code;joinForeignKey:UserCode;References:Code;joinReferences:RoleCode"`
}
```

---

## 4. Protobuf (gRPC) 契约设计 (向下兼容 Java 前端)

为了兼容原有前端的返回结构定义，`identity.proto` 的响应体必须与原先的 JSON 层级完全保持一致。

```protobuf
syntax = "proto3";

package identity;

option go_package = "./pb";

// --- 消息体定义 ---

message GetUserAuthInfoReq {
    string username = 1;
    string tenant_code = 2; // 可选传入
}

message GetUserAuthInfoResp {
    string user_code = 1;
    string password_hash = 2;
    string tenant_code = 3;
    int32  status = 4; // 账号状态，用于 Auth 判断是否被禁用
}

message GetUserInfoReq {
    string user_code = 1;
    string tenant_code = 2;
}

// 嵌套定义 User 实体以匹配原系统 {"user":{...}, "roles":[], "perms":[]}
message UserEntity {
    string user_code = 1;
    string username = 2;
    string avatar = 3;
}

message GetUserInfoResp {
    UserEntity user = 1;
    repeated string roles = 2;
    repeated string perms = 3; // 权限字集合
}

// --- 服务定义 ---

service IdentityService {
    // 【内部校验】提供给 Auth 认证中心专用的内部接口，用于拉取密码 Hash 及基本身份
    rpc GetUserAuthInfo(GetUserAuthInfoReq) returns (GetUserAuthInfoResp);

    // 【用户信息】提供给网关或其他服务拉取用户的全量聚合信息（带上缓存设计）
    rpc GetUserInfo(GetUserInfoReq) returns (GetUserInfoResp);
}
```

## 5. 迁移与开发路径建议
1. **生成框架**：在 `app/identity` 中编写完 `identity.proto` 后，使用 `goctl rpc protoc` 生成骨架代码。
2. **连接数据库与 Redis**：在 `identity.yaml` 中配置 PostgreSQL 与 Redis。
3. **编写数据层**：利用工具（如 `goctl model pg` 或 gorm gen）生成所有对应表结构的 PO 模型，并补齐自定义软删除 Scope。
4. **填充业务**：实现 `GetUserAuthInfo` 与 `GetUserInfo` 链路，重构 Redis 会话写入逻辑。
