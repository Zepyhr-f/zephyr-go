-- ----------------------------
-- 基础设施模块 (Infrastructure) 建表脚本
-- ----------------------------




-- ----------------------------
-- 1. 字典类型表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_dict_type" CASCADE;
CREATE TABLE "zephyr_sys_dict_type" (
  "id"            BIGINT       NOT NULL ,
  "dict_name"     VARCHAR(100) NOT NULL ,
  "dict_type"     VARCHAR(100) NOT NULL ,
  "status"        SMALLINT   NOT NULL DEFAULT 0 ,
  "remark"        VARCHAR(500) DEFAULT NULL ,
  
  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id"),
  UNIQUE ("dict_type", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_dict_type" IS '字典类型表';
COMMENT ON COLUMN "zephyr_sys_dict_type"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_dict_type"."dict_name" IS '字典名称';
COMMENT ON COLUMN "zephyr_sys_dict_type"."dict_type" IS '字典类型 (编码)';
COMMENT ON COLUMN "zephyr_sys_dict_type"."status" IS '状态 (0=正常, 1=停用)';
COMMENT ON COLUMN "zephyr_sys_dict_type"."remark" IS '备注';
COMMENT ON COLUMN "zephyr_sys_dict_type"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_dict_type"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_dict_type"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_dict_type"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_dict_type"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_dict_type"."if_deleted" IS '删除标志 (0=正常, 1=已删除)';


-- ----------------------------
-- 2. 字典数据表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_dict_data" CASCADE;
CREATE TABLE "zephyr_sys_dict_data" (
  "id"            BIGINT       NOT NULL ,
  "dict_sort"     INTEGER          DEFAULT 0 ,
  "dict_label"    VARCHAR(100) NOT NULL ,
  "dict_value"    VARCHAR(100) NOT NULL ,
  "dict_type"     VARCHAR(100) NOT NULL ,
  "css_class"     VARCHAR(100) DEFAULT NULL ,
  "list_class"    VARCHAR(100) DEFAULT NULL ,
  "is_default"    SMALLINT   DEFAULT 0 ,
  "status"        SMALLINT   NOT NULL DEFAULT 0 ,
  "remark"        VARCHAR(500) DEFAULT NULL ,
  
  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_dict_data" IS '字典数据表';
COMMENT ON COLUMN "zephyr_sys_dict_data"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_dict_data"."dict_sort" IS '字典排序';
COMMENT ON COLUMN "zephyr_sys_dict_data"."dict_label" IS '字典标签';
COMMENT ON COLUMN "zephyr_sys_dict_data"."dict_value" IS '字典键值';
COMMENT ON COLUMN "zephyr_sys_dict_data"."dict_type" IS '字典类型';
COMMENT ON COLUMN "zephyr_sys_dict_data"."css_class" IS '样式属性';
COMMENT ON COLUMN "zephyr_sys_dict_data"."list_class" IS '表格回显样式';
COMMENT ON COLUMN "zephyr_sys_dict_data"."is_default" IS '是否默认 (1=是, 0=否)';
COMMENT ON COLUMN "zephyr_sys_dict_data"."status" IS '状态 (0=正常, 1=停用)';
COMMENT ON COLUMN "zephyr_sys_dict_data"."remark" IS '备注';
COMMENT ON COLUMN "zephyr_sys_dict_data"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_dict_data"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_dict_data"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_dict_data"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_dict_data"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_dict_data"."if_deleted" IS '删除标志 (0=正常, 1=已删除)';
CREATE INDEX "idx_dict_type" ON "zephyr_sys_dict_data" ("dict_type");


-- ----------------------------
-- 3. 参数配置表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_config" CASCADE;
CREATE TABLE "zephyr_sys_config" (
  "id"            BIGINT       NOT NULL ,
  "config_name"   VARCHAR(100) NOT NULL ,
  "config_key"    VARCHAR(100) NOT NULL ,
  "config_value"  VARCHAR(500) NOT NULL ,
  "config_type"   SMALLINT   DEFAULT 0 ,
  "remark"        VARCHAR(500) DEFAULT NULL ,
  
  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id"),
  UNIQUE ("config_key", "if_deleted", "tenant_code")
);

COMMENT ON TABLE "zephyr_sys_config" IS '参数配置表';
COMMENT ON COLUMN "zephyr_sys_config"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_config"."config_name" IS '参数名称';
COMMENT ON COLUMN "zephyr_sys_config"."config_key" IS '参数键名';
COMMENT ON COLUMN "zephyr_sys_config"."config_value" IS '参数键值';
COMMENT ON COLUMN "zephyr_sys_config"."config_type" IS '系统内置 (1=是, 0=否)';
COMMENT ON COLUMN "zephyr_sys_config"."remark" IS '备注';
COMMENT ON COLUMN "zephyr_sys_config"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_config"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_config"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_config"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_config"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_config"."if_deleted" IS '删除标志 (0=正常, 1=已删除)';


-- ----------------------------
-- 4. 文件表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_file" CASCADE;
CREATE TABLE "zephyr_sys_file" (
  "id"            BIGINT       NOT NULL ,
  "file_name"     VARCHAR(200) NOT NULL ,
  "file_path"     VARCHAR(500) NOT NULL ,
  "file_url"      VARCHAR(500) NOT NULL ,
  "file_size"     BIGINT       DEFAULT 0 ,
  "file_suffix"   VARCHAR(20)  DEFAULT NULL ,
  "store_type"    VARCHAR(20)  DEFAULT 'local' ,
  
  -- 基础审计与租户字段
  "tenant_code"   VARCHAR(12)  DEFAULT '000000' ,
  "create_user"   VARCHAR(64)  DEFAULT NULL ,
  "create_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"   VARCHAR(64)  DEFAULT NULL ,
  "update_time"   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"    SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_file" IS '文件记录表';
COMMENT ON COLUMN "zephyr_sys_file"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_file"."file_name" IS '原始文件名';
COMMENT ON COLUMN "zephyr_sys_file"."file_path" IS '存储相对路径';
COMMENT ON COLUMN "zephyr_sys_file"."file_url" IS '访问全路径';
COMMENT ON COLUMN "zephyr_sys_file"."file_size" IS '文件大小 (Byte)';
COMMENT ON COLUMN "zephyr_sys_file"."file_suffix" IS '扩展名';
COMMENT ON COLUMN "zephyr_sys_file"."store_type" IS '存储方案';
COMMENT ON COLUMN "zephyr_sys_file"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_file"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_file"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_file"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_file"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_file"."if_deleted" IS '删除标志 (0=正常, 1=已删除)';


-- ----------------------------
-- 5. 通知公告表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_notice" CASCADE;
CREATE TABLE "zephyr_sys_notice" (
  "id"             BIGINT       NOT NULL ,
  "notice_title"   VARCHAR(200) NOT NULL ,
  "notice_type"    SMALLINT   NOT NULL DEFAULT 1 ,
  "notice_content" TEXT     NOT NULL ,
  "status"         SMALLINT   NOT NULL DEFAULT 0 ,
  "remark"         VARCHAR(500) DEFAULT NULL ,
  
  -- 基础审计与租户字段
  "tenant_code"    VARCHAR(12)  DEFAULT '000000' ,
  "create_user"    VARCHAR(64)  DEFAULT NULL ,
  "create_time"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"    VARCHAR(64)  DEFAULT NULL ,
  "update_time"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"     SMALLINT   NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_notice" IS '通知公告表';
COMMENT ON COLUMN "zephyr_sys_notice"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_notice"."notice_title" IS '公告标题';
COMMENT ON COLUMN "zephyr_sys_notice"."notice_type" IS '公告类型 (1=系统通知, 2=业务提醒)';
COMMENT ON COLUMN "zephyr_sys_notice"."notice_content" IS '公告内容';
COMMENT ON COLUMN "zephyr_sys_notice"."status" IS '公告状态 (0=草稿, 1=已发布, 2=已撤退)';
COMMENT ON COLUMN "zephyr_sys_notice"."remark" IS '备注';
COMMENT ON COLUMN "zephyr_sys_notice"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_notice"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_notice"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_notice"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_notice"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_notice"."if_deleted" IS '删除标志 (0=正常, 1=已删除)';



