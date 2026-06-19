-- ============================================================================
-- 任务调度模块建表脚本
-- ----------------------------
-- 1. 定时任务调度表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_job" CASCADE;
CREATE TABLE "zephyr_sys_job" (
  "id"              BIGINT        NOT NULL ,
  "job_name"        VARCHAR(64)   DEFAULT '' ,
  "job_group"       VARCHAR(64)   DEFAULT 'DEFAULT' ,
  "invoke_target"   VARCHAR(500)  NOT NULL ,
  "cron_expression" VARCHAR(255)  DEFAULT '' ,
  "misfire_policy"  SMALLINT    DEFAULT 3 ,
  "concurrent"      SMALLINT    DEFAULT 1 ,
  "status"          SMALLINT    DEFAULT 0 ,
  "remark"          VARCHAR(500)  DEFAULT '' ,
  
  -- 基础审计与租户字段
  "tenant_code"     VARCHAR(12)   DEFAULT '000000' ,
  "create_user"     VARCHAR(64)   DEFAULT NULL ,
  "create_time"     TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ,
  "update_user"     VARCHAR(64)   DEFAULT NULL ,
  "update_time"     TIMESTAMP      DEFAULT CURRENT_TIMESTAMP  ,
  "if_deleted"      SMALLINT    NOT NULL DEFAULT 0 ,
  
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_job" IS '定时任务调度表';
COMMENT ON COLUMN "zephyr_sys_job"."id" IS '任务ID';
COMMENT ON COLUMN "zephyr_sys_job"."job_name" IS '任务名称';
COMMENT ON COLUMN "zephyr_sys_job"."job_group" IS '任务组名';
COMMENT ON COLUMN "zephyr_sys_job"."invoke_target" IS '调用目标字符串';
COMMENT ON COLUMN "zephyr_sys_job"."cron_expression" IS 'cron执行表达式';
COMMENT ON COLUMN "zephyr_sys_job"."misfire_policy" IS '计划执行策略（1立即执行 2执行一次 3放弃执行）';
COMMENT ON COLUMN "zephyr_sys_job"."concurrent" IS '是否并发执行（0禁止 1允许）';
COMMENT ON COLUMN "zephyr_sys_job"."status" IS '状态（0正常 1暂停）';
COMMENT ON COLUMN "zephyr_sys_job"."remark" IS '备注信息';
COMMENT ON COLUMN "zephyr_sys_job"."tenant_code" IS '租户编码';
COMMENT ON COLUMN "zephyr_sys_job"."create_user" IS '创建者编码';
COMMENT ON COLUMN "zephyr_sys_job"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_job"."update_user" IS '更新者编码';
COMMENT ON COLUMN "zephyr_sys_job"."update_time" IS '更新时间';
COMMENT ON COLUMN "zephyr_sys_job"."if_deleted" IS '删除标志（0=正常 1=已删除）';


-- ----------------------------
-- 2. 定时任务调度日志表
-- ----------------------------
DROP TABLE IF EXISTS "zephyr_sys_job_log" CASCADE;
CREATE TABLE "zephyr_sys_job_log" (
  "id"              BIGINT        NOT NULL ,
  "job_name"        VARCHAR(64)   NOT NULL ,
  "job_group"       VARCHAR(64)   NOT NULL ,
  "invoke_target"   VARCHAR(500)  NOT NULL ,
  "job_message"     VARCHAR(500)  ,
  "status"          SMALLINT    DEFAULT 0 ,
  "exception_info"  VARCHAR(2000) DEFAULT '' ,
  "create_time"     TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ,
  "cost_time"       BIGINT        DEFAULT 0 ,
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "zephyr_sys_job_log" IS '定时任务调度日志表';
COMMENT ON COLUMN "zephyr_sys_job_log"."id" IS '记录ID';
COMMENT ON COLUMN "zephyr_sys_job_log"."job_name" IS '任务名称';
COMMENT ON COLUMN "zephyr_sys_job_log"."job_group" IS '任务组名';
COMMENT ON COLUMN "zephyr_sys_job_log"."invoke_target" IS '调用目标字符串';
COMMENT ON COLUMN "zephyr_sys_job_log"."job_message" IS '日志信息';
COMMENT ON COLUMN "zephyr_sys_job_log"."status" IS '执行状态（0正常 1失败）';
COMMENT ON COLUMN "zephyr_sys_job_log"."exception_info" IS '异常信息';
COMMENT ON COLUMN "zephyr_sys_job_log"."create_time" IS '创建时间';
COMMENT ON COLUMN "zephyr_sys_job_log"."cost_time" IS '消耗时间(ms)';

