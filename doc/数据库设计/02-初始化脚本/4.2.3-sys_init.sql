-- ----------------------------
-- 禁用外键检查
-- ----------------------------


-- ----------------------------
-- 清空数据
-- ----------------------------
DELETE FROM "zephyr_sys_user_role";
DELETE FROM "zephyr_sys_user_post";
DELETE FROM "zephyr_sys_dict";
DELETE FROM "zephyr_sys_user";
DELETE FROM "zephyr_sys_dept";
DELETE FROM "zephyr_sys_role";
DELETE FROM "zephyr_sys_post";

-- ----------------------------
-- 启用外键检查
-- ----------------------------


-- ----------------------------
-- 1. 字典数据
-- ----------------------------
INSERT INTO "zephyr_sys_dict" ("id", "parent_id", "code", "dict_key", "dict_value", "sort", "remark", "status", "tenant_code", "create_user") VALUES
(1, 0, 'sex', '0', '男', 1, '性别', 1, '000000', 1),
(2, 0, 'sex', '1', '女', 2, '性别', 1, '000000', 1),
(3, 0, 'user_status', '1', '正常', 1, '用户状态', 1, '000000', 1),
(4, 0, 'user_status', '0', '停用', 2, '用户状态', 1, '000000', 1);

-- ----------------------------
-- 2. 部门数据
-- ----------------------------
INSERT INTO "zephyr_sys_dept" ("id", "code", "parent_code", "dept_name", "full_name", "order_num", "status", "tenant_code", "create_user") VALUES
(1, 'ZEPHYR', '0', '总公司', '科技有限公司总公司', 1, 1, '000000', 1),
(2, 'RD', 'ZEPHYR', '技术部', '科技有限公司技术部', 1, 1, '000000', 1),
(3, 'MKT', 'ZEPHYR', '市场部', '科技有限公司市场部', 2, 1, '000000', 1);

-- ----------------------------
-- 3. 角色数据
-- ----------------------------
INSERT INTO "zephyr_sys_role" ("id", "code", "role_name", "order_num", "status", "tenant_code", "create_user") VALUES
(1, 'admin', '超级管理员', 1, 1, '000000', 1),
(2, 'manager', '部门经理', 2, 1, '000000', 1),
(3, 'employee', '普通员工', 3, 1, '000000', 1);

-- ----------------------------
-- 4. 岗位数据
-- ----------------------------
INSERT INTO "zephyr_sys_post" ("id", "code", "post_name", "order_num", "status", "tenant_code", "create_user") VALUES
(1, 'CEO', '首席执行官', 1, 1, '000000', 1),
(2, 'CTO', '技术总监', 2, 1, '000000', 1),
(3, 'DEV', '开发工程师', 3, 1, '000000', 1);

-- ----------------------------
-- 5. 用户数据
-- ----------------------------
INSERT INTO "zephyr_sys_user" ("id", "code", "nick_name", "real_name", "password", "email", "phone", "sex", "dept_code", "post_code", "status", "tenant_code", "create_user") VALUES
(1, 'U10001', 'admin', '张管理员', '$2a$10$V4cAPzRXVt948dqR6cAIFOlLh29/RHEDy7BjbfYYm0eiXRGIx3ZKy', 'admin@company.com', '13800138001', 0, 'ZEPHYR', 'CEO', 1, '000000', 1),
(2, 'U10002', 'zhangsan', '张三', '$2a$10$V4cAPzRXVt948dqR6cAIFOlLh29/RHEDy7BjbfYYm0eiXRGIx3ZKy', 'zhangsan@company.com', '13800138002', 0, 'RD', 'CTO', 1, '000000', 1),
(3, 'U10003', 'lisi', '李四', '$2a$10$V4cAPzRXVt948dqR6cAIFOlLh29/RHEDy7BjbfYYm0eiXRGIx3ZKy', 'lisi@company.com', '13800138003', 0, 'RD', 'DEV', 1, '000000', 1);

-- ----------------------------
-- 6. 关联数据
-- ----------------------------
INSERT INTO "zephyr_sys_user_role" ("user_code", "role_code") VALUES
('U10001', 'admin'),
('U10002', 'manager'),
('U10003', 'employee');

INSERT INTO "zephyr_sys_user_post" ("user_code", "post_code") VALUES
('U10001', 'CEO'),
('U10002', 'CTO'),
('U10003', 'DEV');