-- ============================================================================
-- 安全审计日志建表脚本
-- 遵循规范：zephyr-doc/05-开发规范/00-建表规范.md
-- ============================================================================

-- ----------------------------
-- 1. 登录日志表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_login_log" CASCADE;
CREATE TABLE "zephyr_sys_login_log" (
  -- 主键
  "id"             BIGINT       NOT NULL ,

  -- 业务字段
  "username"       VARCHAR(50)  DEFAULT '' ,
  "ipaddr"         VARCHAR(128) DEFAULT '' ,
  "login_location" VARCHAR(255) DEFAULT '' ,
  "browser"        VARCHAR(50)  DEFAULT '' ,
  "os"             VARCHAR(50)  DEFAULT '' ,
  "status"         SMALLINT   DEFAULT 0 ,
  "msg"            VARCHAR(255) DEFAULT '' ,
  "login_time"     TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,

  -- 基础审计与租户字段
  "tenant_code"    VARCHAR(12)  DEFAULT '000000' ,
  "create_user"    VARCHAR(64)  DEFAULT NULL ,
  "create_time"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ,
  "update_user"    VARCHAR(64)  DEFAULT NULL ,
  "update_time"    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"     SMALLINT   NOT NULL DEFAULT 0 ,

  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_login_log" IS '系统访问记录表';
COMMENT ON COLUMN "zephyr_sys_login_log"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_login_log"."username" IS '登录账号';
COMMENT ON COLUMN "zephyr_sys_login_log"."ipaddr" IS '登录IP地址';
COMMENT ON COLUMN "zephyr_sys_login_log"."login_location" IS '登录地点（IP解析）';
COMMENT ON COLUMN "zephyr_sys_login_log"."browser" IS '浏览器类型';
COMMENT ON COLUMN "zephyr_sys_login_log"."os" IS '操作系统';
COMMENT ON COLUMN "zephyr_sys_login_log"."status" IS '登录状态（0=成功 1=失败）';
COMMENT ON COLUMN "zephyr_sys_login_log"."msg" IS '提示消息';
COMMENT ON COLUMN "zephyr_sys_login_log"."login_time" IS '登录时间';
COMMENT ON COLUMN "zephyr_sys_login_log"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_login_log"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_login_log"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_login_log"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_login_log"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_login_log"."if_deleted" IS '删除标识（0=正常 1=已删除）';
CREATE INDEX "idx_username" ON "zephyr_sys_login_log" ("username");
CREATE INDEX "idx_login_time" ON "zephyr_sys_login_log" ("login_time");


-- ----------------------------
-- 2. 操作日志表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_oper_log" CASCADE;
CREATE TABLE "zephyr_sys_oper_log" (
  -- 主键
  "id"             BIGINT        NOT NULL ,

  -- 业务字段
  "title"          VARCHAR(50)   DEFAULT '' ,
  "business_type"  SMALLINT    DEFAULT 0 ,
  "method"         VARCHAR(100)  DEFAULT '' ,
  "request_method" VARCHAR(10)   DEFAULT '' ,
  "operator_type"  SMALLINT    DEFAULT 0 ,
  "oper_name"      VARCHAR(50)   DEFAULT '' ,
  "dept_name"      VARCHAR(50)   DEFAULT '' ,
  "oper_url"       VARCHAR(255)  DEFAULT '' ,
  "oper_ip"        VARCHAR(128)  DEFAULT '' ,
  "oper_location"  VARCHAR(255)  DEFAULT '' ,
  "oper_param"     VARCHAR(2000) DEFAULT '' ,
  "json_result"    VARCHAR(2000) DEFAULT '' ,
  "status"         SMALLINT    DEFAULT 0 ,
  "error_msg"      VARCHAR(2000) DEFAULT '' ,
  "oper_time"      TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ,
  "cost_time"      BIGINT        DEFAULT 0 ,

  -- 基础审计与租户字段
  "tenant_code"    VARCHAR(12)   DEFAULT '000000' ,
  "create_user"    VARCHAR(64)   DEFAULT NULL ,
  "create_time"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ,
  "update_user"    VARCHAR(64)   DEFAULT NULL ,
  "update_time"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"     SMALLINT    NOT NULL DEFAULT 0 ,

  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_oper_log" IS '操作日志记录表';
COMMENT ON COLUMN "zephyr_sys_oper_log"."id" IS '主键ID';
COMMENT ON COLUMN "zephyr_sys_oper_log"."title" IS '系统模块名称';
COMMENT ON COLUMN "zephyr_sys_oper_log"."business_type" IS '业务类型（0=其它 1=新增 2=修改 3=删除 4=授权 5=导出 6=导入 7=强退 8=生成代码 9=清空数据）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."method" IS '方法名称';
COMMENT ON COLUMN "zephyr_sys_oper_log"."request_method" IS '请求方式（POST/PUT/DELETE等）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."operator_type" IS '操作类别（0=其它 1=后台用户 2=手机端用户）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_name" IS '操作人员账号';
COMMENT ON COLUMN "zephyr_sys_oper_log"."dept_name" IS '部门名称';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_url" IS '请求URL';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_ip" IS '主机地址';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_location" IS '操作地点（IP解析）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_param" IS '请求参数';
COMMENT ON COLUMN "zephyr_sys_oper_log"."json_result" IS '返回参数';
COMMENT ON COLUMN "zephyr_sys_oper_log"."status" IS '操作状态（0=正常 1=异常）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."error_msg" IS '错误消息';
COMMENT ON COLUMN "zephyr_sys_oper_log"."oper_time" IS '操作时间';
COMMENT ON COLUMN "zephyr_sys_oper_log"."cost_time" IS '消耗时间（ms）';
COMMENT ON COLUMN "zephyr_sys_oper_log"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_oper_log"."create_user" IS '创建人编码';
COMMENT ON COLUMN "zephyr_sys_oper_log"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_oper_log"."update_user" IS '更新人编码';
COMMENT ON COLUMN "zephyr_sys_oper_log"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_oper_log"."if_deleted" IS '删除标识（0=正常 1=已删除）';
CREATE INDEX "idx_oper_name" ON "zephyr_sys_oper_log" ("oper_name");
CREATE INDEX "idx_oper_time" ON "zephyr_sys_oper_log" ("oper_time");
CREATE INDEX "idx_status" ON "zephyr_sys_oper_log" ("status");

