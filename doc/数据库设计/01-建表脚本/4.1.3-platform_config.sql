-- ----------------------------
-- 1. 字典表 (旧版，建议参考 infrastructure_module.sql 中的双表设计)
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_dict" CASCADE;
CREATE TABLE "zephyr_sys_dict" (
  "id"            BIGINT       NOT NULL ,
  "parent_id"     BIGINT       DEFAULT 0 ,
  "code"          VARCHAR(255) DEFAULT NULL ,
  "dict_key"      VARCHAR(255) DEFAULT NULL ,
  "dict_value"    VARCHAR(255) DEFAULT NULL ,
  "sort"          INTEGER          DEFAULT 0 ,
  "remark"        VARCHAR(255) DEFAULT NULL ,
  "status"        SMALLINT   DEFAULT 1 ,
  
  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_dict" IS '字典表';
COMMENT ON COLUMN "zephyr_sys_dict"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_dict"."parent_id" IS '父主键';
COMMENT ON COLUMN "zephyr_sys_dict"."code" IS '字典码';
COMMENT ON COLUMN "zephyr_sys_dict"."dict_key" IS '字典值';
COMMENT ON COLUMN "zephyr_sys_dict"."dict_value" IS '字典名称';
COMMENT ON COLUMN "zephyr_sys_dict"."sort" IS '排序';
COMMENT ON COLUMN "zephyr_sys_dict"."remark" IS '字典备注';
COMMENT ON COLUMN "zephyr_sys_dict"."status" IS '状态 (1-正常, 0-禁用)';
COMMENT ON COLUMN "zephyr_sys_dict"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_dict"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_dict"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_dict"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_dict"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_dict"."if_deleted" IS '删除标识（0=正常 1=已删除）';


-- ----------------------------
-- 2. 岗位表 (已在 rbac_core.sql 中定义，此处保留并对齐规范)
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_post_bak" CASCADE;
CREATE TABLE "zephyr_sys_post_bak" (
  "id"            BIGINT       NOT NULL ,
  "category"      INTEGER          DEFAULT NULL ,
  "post_code"     VARCHAR(64)  NOT NULL ,
  "post_name"     VARCHAR(64)  NOT NULL ,
  "sort"          INTEGER          DEFAULT 0 ,
  "remark"        VARCHAR(255) DEFAULT NULL ,
  "status"        SMALLINT   DEFAULT 1 ,

  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,

  PRIMARY KEY ("id"),
  UNIQUE ("post_code", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_post_bak" IS '岗位表(备用)';
COMMENT ON COLUMN "zephyr_sys_post_bak"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_post_bak"."category" IS '岗位类型';
COMMENT ON COLUMN "zephyr_sys_post_bak"."post_code" IS '岗位编号';
COMMENT ON COLUMN "zephyr_sys_post_bak"."post_name" IS '岗位名称';
COMMENT ON COLUMN "zephyr_sys_post_bak"."sort" IS '岗位排序';
COMMENT ON COLUMN "zephyr_sys_post_bak"."remark" IS '岗位描述';
COMMENT ON COLUMN "zephyr_sys_post_bak"."status" IS '状态 (1-正常, 0-禁用)';
COMMENT ON COLUMN "zephyr_sys_post_bak"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_post_bak"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_post_bak"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_post_bak"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_post_bak"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_post_bak"."if_deleted" IS '删除标识（0=正常 1=已删除）';

