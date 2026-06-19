-- ----------------------------
-- 删除表（按依赖关系逆序）
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_role_dept" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_role_menu" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_user_post" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_user_role" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_menu" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_post" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_role" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_user" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_dept" CASCADE;
DROP TABLE IF EXISTS "zephyr_sys_tenant" CASCADE;

-- ----------------------------
-- 1. 租户信息表（系统级，tenant_code 固定为 000000）
--    继承基类：BaseEntity + code（业务唯一标识）
-- ----------------------------
CREATE TABLE "zephyr_sys_tenant" (
  "id"              BIGINT       NOT NULL ,
  "code"            VARCHAR(12)  NOT NULL ,
  "tenant_name"     VARCHAR(100) NOT NULL ,
  "contact_user"    VARCHAR(50)  DEFAULT NULL ,
  "contact_phone"   VARCHAR(20)  DEFAULT NULL ,
  "status"          SMALLINT   DEFAULT 1 ,
  "expire_time"     TIMESTAMP     DEFAULT NULL ,
  "account_count"   INTEGER          DEFAULT -1 ,
  "remark"          VARCHAR(500) DEFAULT NULL ,
  "tenant_code"     VARCHAR(12)  DEFAULT '000000' ,
  "create_user"     VARCHAR(64)  DEFAULT NULL ,
  "create_time"     TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"     VARCHAR(64)  DEFAULT NULL ,
  "update_time"     TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"      SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted")
);

COMMENT ON TABLE "zephyr_sys_tenant" IS '租户信息表';
COMMENT ON COLUMN "zephyr_sys_tenant"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_tenant"."code" IS '租户编号 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_tenant"."tenant_name" IS '租户名称/公司名称';
COMMENT ON COLUMN "zephyr_sys_tenant"."contact_user" IS '联系人';
COMMENT ON COLUMN "zephyr_sys_tenant"."contact_phone" IS '联系电话';
COMMENT ON COLUMN "zephyr_sys_tenant"."status" IS '租户状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_tenant"."expire_time" IS '授权过期时间';
COMMENT ON COLUMN "zephyr_sys_tenant"."account_count" IS '账号额度 (-1表示无限制)';
COMMENT ON COLUMN "zephyr_sys_tenant"."remark" IS '备注';
COMMENT ON COLUMN "zephyr_sys_tenant"."tenant_code" IS '租户编码 (系统级固定000000)';
COMMENT ON COLUMN "zephyr_sys_tenant"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_tenant"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_tenant"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_tenant"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_tenant"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 2. 部门表
--    继承基类：TreeEntity (BaseEntity + code + parent_code + leaf)
-- ----------------------------
CREATE TABLE "zephyr_sys_dept" (
  "id"          BIGINT       NOT NULL ,
  "code"        VARCHAR(64)  NOT NULL ,
  "parent_code" VARCHAR(64)  DEFAULT NULL ,
  "leaf"        SMALLINT   DEFAULT 1 ,
  "dept_name"   VARCHAR(50)  NOT NULL ,
  "full_name"   VARCHAR(255) DEFAULT NULL ,
  "order_num"   INTEGER          DEFAULT 0 ,
  "status"      SMALLINT   DEFAULT 1 ,
  "tenant_code" VARCHAR(12)  DEFAULT '000000' ,
  "create_user" VARCHAR(64)  DEFAULT NULL ,
  "create_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user" VARCHAR(64)  DEFAULT NULL ,
  "update_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"  SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted", "tenant_code")
);
CREATE INDEX "idx_parent_code" ON "zephyr_sys_dept" ("parent_code");

COMMENT ON TABLE "zephyr_sys_dept" IS '部门表';
COMMENT ON COLUMN "zephyr_sys_dept"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_dept"."code" IS '部门编码 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_dept"."parent_code" IS '父部门编码 (NULL表示根节点)';
COMMENT ON COLUMN "zephyr_sys_dept"."leaf" IS '是否叶子节点 (1=是 0=否)';
COMMENT ON COLUMN "zephyr_sys_dept"."dept_name" IS '部门名称';
COMMENT ON COLUMN "zephyr_sys_dept"."full_name" IS '部门全称';
COMMENT ON COLUMN "zephyr_sys_dept"."order_num" IS '显示顺序';
COMMENT ON COLUMN "zephyr_sys_dept"."status" IS '部门状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_dept"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_dept"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_dept"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_dept"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_dept"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_dept"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 3. 岗位信息表
--    继承基类：CodeEntity (BaseEntity + code)
-- ----------------------------
CREATE TABLE "zephyr_sys_post" (
  "id"          BIGINT       NOT NULL ,
  "code"        VARCHAR(64)  NOT NULL ,
  "post_name"   VARCHAR(50)  NOT NULL ,
  "order_num"   INTEGER          DEFAULT 0 ,
  "status"      SMALLINT   DEFAULT 1 ,
  "tenant_code" VARCHAR(12)  DEFAULT '000000' ,
  "create_user" VARCHAR(64)  DEFAULT NULL ,
  "create_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user" VARCHAR(64)  DEFAULT NULL ,
  "update_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"  SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_post" IS '岗位信息表';
COMMENT ON COLUMN "zephyr_sys_post"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_post"."code" IS '岗位编码 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_post"."post_name" IS '岗位名称';
COMMENT ON COLUMN "zephyr_sys_post"."order_num" IS '显示顺序';
COMMENT ON COLUMN "zephyr_sys_post"."status" IS '状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_post"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_post"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_post"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_post"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_post"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_post"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 4. 用户表
--    继承基类：CodeEntity (BaseEntity + code)
-- ----------------------------
CREATE TABLE "zephyr_sys_user" (
  "id"          BIGINT       NOT NULL ,
  "code"        VARCHAR(64)  NOT NULL ,
  "nick_name"   VARCHAR(50)  NOT NULL ,
  "real_name"   VARCHAR(50)  DEFAULT NULL ,
  "password"    VARCHAR(100) NOT NULL ,
  "avatar"      VARCHAR(255) DEFAULT NULL ,
  "email"       VARCHAR(100) DEFAULT NULL ,
  "phone"       VARCHAR(20)  DEFAULT NULL ,
  "sex"         SMALLINT   DEFAULT 0 ,
  "birthday"    DATE         DEFAULT NULL ,
  "user_type"   SMALLINT   DEFAULT 0 ,
  "status"      SMALLINT   DEFAULT 1 ,
  "dept_code"   VARCHAR(64)  DEFAULT NULL ,
  "post_code"   VARCHAR(64)  DEFAULT NULL ,
  "tenant_code" VARCHAR(12)  DEFAULT '000000' ,
  "create_user" VARCHAR(64)  DEFAULT NULL ,
  "create_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user" VARCHAR(64)  DEFAULT NULL ,
  "update_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"  SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_user" IS '用户信息表';
COMMENT ON COLUMN "zephyr_sys_user"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_user"."code" IS '员工编码 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_user"."nick_name" IS '昵称';
COMMENT ON COLUMN "zephyr_sys_user"."real_name" IS '真实姓名';
COMMENT ON COLUMN "zephyr_sys_user"."password" IS '密码';
COMMENT ON COLUMN "zephyr_sys_user"."avatar" IS '头像';
COMMENT ON COLUMN "zephyr_sys_user"."email" IS '邮箱';
COMMENT ON COLUMN "zephyr_sys_user"."phone" IS '手机号';
COMMENT ON COLUMN "zephyr_sys_user"."sex" IS '性别 (0=男 1=女)';
COMMENT ON COLUMN "zephyr_sys_user"."birthday" IS '出生日期';
COMMENT ON COLUMN "zephyr_sys_user"."user_type" IS '用户类型 (0=员工 1=管理员 2=系统)';
COMMENT ON COLUMN "zephyr_sys_user"."status" IS '帐号状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_user"."dept_code" IS '所属部门编码 (逻辑关联 sys_dept.code)';
COMMENT ON COLUMN "zephyr_sys_user"."post_code" IS '所属岗位编码 (逻辑关联 sys_post.code)';
COMMENT ON COLUMN "zephyr_sys_user"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_user"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_user"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_user"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_user"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_user"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 5. 角色表
--    继承基类：CodeEntity (BaseEntity + code)
-- ----------------------------
CREATE TABLE "zephyr_sys_role" (
  "id"          BIGINT       NOT NULL ,
  "code"        VARCHAR(64)  NOT NULL ,
  "role_name"   VARCHAR(30)  NOT NULL ,
  "order_num"   INTEGER          DEFAULT 0 ,
  "status"      SMALLINT   DEFAULT 1 ,
  "tenant_code" VARCHAR(12)  DEFAULT '000000' ,
  "create_user" VARCHAR(64)  DEFAULT NULL ,
  "create_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user" VARCHAR(64)  DEFAULT NULL ,
  "update_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"  SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_role" IS '角色信息表';
COMMENT ON COLUMN "zephyr_sys_role"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_role"."code" IS '角色编码 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_role"."role_name" IS '角色名称';
COMMENT ON COLUMN "zephyr_sys_role"."order_num" IS '显示顺序';
COMMENT ON COLUMN "zephyr_sys_role"."status" IS '角色状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_role"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_role"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_role"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_role"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_role"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_role"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 6. 菜单表
--    继承基类：TreeEntity (BaseEntity + code + parent_code + leaf)
-- ----------------------------
CREATE TABLE "zephyr_sys_menu" (
  "id"          BIGINT       NOT NULL ,
  "code"        VARCHAR(64)  NOT NULL ,
  "parent_code" VARCHAR(64)  DEFAULT NULL ,
  "leaf"        SMALLINT   DEFAULT 1 ,
  "menu_name"   VARCHAR(30)  NOT NULL ,
  "menu_type"   CHAR(1)      DEFAULT 'M' ,
  "path"        VARCHAR(200) DEFAULT '' ,
  "component"   VARCHAR(255) DEFAULT NULL ,
  "perms"       VARCHAR(100) DEFAULT NULL ,
  "icon"        VARCHAR(100) DEFAULT '#' ,
  "order_num"   INTEGER          DEFAULT 0 ,
  "status"      SMALLINT   DEFAULT 1 ,
  "tenant_code" VARCHAR(12)  DEFAULT '000000' ,
  "create_user" VARCHAR(64)  DEFAULT NULL ,
  "create_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user" VARCHAR(64)  DEFAULT NULL ,
  "update_time" TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"  SMALLINT   NOT NULL DEFAULT 0 ,
  PRIMARY KEY ("id"),
  UNIQUE ("code", "if_deleted", "tenant_code")
);
CREATE INDEX "idx_menu_parent_code" ON "zephyr_sys_menu" ("parent_code");

COMMENT ON TABLE "zephyr_sys_menu" IS '菜单与权限规则表';
COMMENT ON COLUMN "zephyr_sys_menu"."id" IS '主键ID (雪花算法)';
COMMENT ON COLUMN "zephyr_sys_menu"."code" IS '菜单编码 (业务唯一标识)';
COMMENT ON COLUMN "zephyr_sys_menu"."parent_code" IS '父菜单编码 (NULL表示根节点)';
COMMENT ON COLUMN "zephyr_sys_menu"."leaf" IS '是否叶子节点 (1=是 0=否)';
COMMENT ON COLUMN "zephyr_sys_menu"."menu_name" IS '菜单名称';
COMMENT ON COLUMN "zephyr_sys_menu"."menu_type" IS '菜单类型 (M=目录 C=菜单 F=按钮/API)';
COMMENT ON COLUMN "zephyr_sys_menu"."path" IS '路由地址';
COMMENT ON COLUMN "zephyr_sys_menu"."component" IS '组件路径';
COMMENT ON COLUMN "zephyr_sys_menu"."perms" IS '权限标识';
COMMENT ON COLUMN "zephyr_sys_menu"."icon" IS '菜单图标';
COMMENT ON COLUMN "zephyr_sys_menu"."order_num" IS '显示顺序';
COMMENT ON COLUMN "zephyr_sys_menu"."status" IS '菜单状态 (1=正常 0=停用)';
COMMENT ON COLUMN "zephyr_sys_menu"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_menu"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_menu"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_menu"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_menu"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_menu"."if_deleted" IS '删除标识 (0=正常 1=已删除)';


-- ----------------------------
-- 7. 关联关系表 (增加索引，无物理外键)
--    中间表不继承 BaseEntity，使用 xxx_code 关联业务表的 code 字段
-- ----------------------------
CREATE TABLE "zephyr_sys_user_role" (
  "user_code" VARCHAR(64) NOT NULL ,
  "role_code" VARCHAR(64) NOT NULL ,
  PRIMARY KEY ("user_code", "role_code")
);
CREATE INDEX "idx_role_code" ON "zephyr_sys_user_role" ("role_code");

COMMENT ON TABLE "zephyr_sys_user_role" IS '用户和角色关联表';
COMMENT ON COLUMN "zephyr_sys_user_role"."user_code" IS '用户编码 (关联 sys_user.code)';
COMMENT ON COLUMN "zephyr_sys_user_role"."role_code" IS '角色编码 (关联 sys_role.code)';


CREATE TABLE "zephyr_sys_user_post" (
  "user_code" VARCHAR(64) NOT NULL ,
  "post_code" VARCHAR(64) NOT NULL ,
  PRIMARY KEY ("user_code", "post_code")
);
CREATE INDEX "idx_post_code" ON "zephyr_sys_user_post" ("post_code");

COMMENT ON TABLE "zephyr_sys_user_post" IS '用户与岗位关联表';
COMMENT ON COLUMN "zephyr_sys_user_post"."user_code" IS '用户编码 (关联 sys_user.code)';
COMMENT ON COLUMN "zephyr_sys_user_post"."post_code" IS '岗位编码 (关联 sys_post.code)';


CREATE TABLE "zephyr_sys_role_menu" (
  "role_code" VARCHAR(64) NOT NULL ,
  "menu_code" VARCHAR(64) NOT NULL ,
  PRIMARY KEY ("role_code", "menu_code")
);
CREATE INDEX "idx_menu_code" ON "zephyr_sys_role_menu" ("menu_code");

COMMENT ON TABLE "zephyr_sys_role_menu" IS '角色和菜单关联表';
COMMENT ON COLUMN "zephyr_sys_role_menu"."role_code" IS '角色编码 (关联 sys_role.code)';
COMMENT ON COLUMN "zephyr_sys_role_menu"."menu_code" IS '菜单编码 (关联 sys_menu.code)';


CREATE TABLE "zephyr_sys_role_dept" (
  "role_code" VARCHAR(64) NOT NULL ,
  "dept_code" VARCHAR(64) NOT NULL ,
  PRIMARY KEY ("role_code", "dept_code")
);
CREATE INDEX "idx_dept_code" ON "zephyr_sys_role_dept" ("dept_code");

COMMENT ON TABLE "zephyr_sys_role_dept" IS '角色和部门关联表';
COMMENT ON COLUMN "zephyr_sys_role_dept"."role_code" IS '角色编码 (关联 sys_role.code)';
COMMENT ON COLUMN "zephyr_sys_role_dept"."dept_code" IS '部门编码 (关联 sys_dept.code)';

