-- ----------------------------
-- 禁用外键检查
-- ----------------------------


-- ----------------------------
-- 清空数据（按依赖关系逆序删除）
-- ----------------------------
DELETE FROM "zephyr_sys_role_dept";
DELETE FROM "zephyr_sys_role_menu";
DELETE FROM "zephyr_sys_user_role";
DELETE FROM "zephyr_sys_menu";
DELETE FROM "zephyr_sys_role";
DELETE FROM "zephyr_sys_user";
DELETE FROM "zephyr_sys_dept";
DELETE FROM "zephyr_sys_tenant";

-- ----------------------------
-- 启用外键检查
-- ----------------------------


-- ===============================
-- 0. 租户信息
-- ===============================
INSERT INTO "zephyr_sys_tenant" (id, code, tenant_name, status, create_user) VALUES
(1, '000000', '默认租户', 1, 1);

-- ===============================
-- 1. 部门
-- ===============================
INSERT INTO "zephyr_sys_dept" (id, code, parent_code, dept_name, full_name, order_num, status, tenant_code, create_user) VALUES
(100, 'ZEPHYR', '0', 'Zephyr集团', 'Zephyr集团总部', 1, 1, '000000', 1),
(101, 'ZEPHYR_CEO', 'ZEPHYR', '总经办', 'Zephyr集团总经办', 1, 1, '000000', 1),
(102, 'ZEPHYR_RD', 'ZEPHYR', '研发部', 'Zephyr集团研发部', 2, 1, '000000', 1);

-- ===============================
-- 2. 用户
-- ===============================
-- 密码均为: 123456 (BCrypt加密)
INSERT INTO "zephyr_sys_user" (id, code, nick_name, real_name, password, email, phone, sex, status, dept_code, tenant_code, create_user) VALUES
(1, 'admin', 'admin', '超级管理员', '$2a$10$V4cAPzRXVt948dqR6cAIFOlLh29/RHEDy7BjbfYYm0eiXRGIx3ZKy', 'admin@zephyr.com', '15800000000', 0, 1, 'ZEPHYR_CEO', '000000', 1),
(2, 'ry', 'ry', '若依', '$2a$10$V4cAPzRXVt948dqR6cAIFOlLh29/RHEDy7BjbfYYm0eiXRGIx3ZKy', 'ry@zephyr.com', '15800000001', 1, 1, 'ZEPHYR_RD', '000000', 1);

-- ===============================
-- 3. 角色
-- ===============================
INSERT INTO "zephyr_sys_role" (id, code, role_name, order_num, status, tenant_code, create_user) VALUES
(1, 'admin', '超级管理员', 1, 1, '000000', 1),
(2, 'common', '普通角色', 2, 1, '000000', 1);

-- ===============================
-- 4. 用户-角色关联 (使用编码关联)
-- ===============================
INSERT INTO "zephyr_sys_user_role" (user_code, role_code) VALUES
('admin', 'admin'),
('ry', 'common');

INSERT INTO "zephyr_sys_menu" (id, code, parent_code, menu_name, menu_type, path, component, icon, perms, order_num, status, tenant_code, create_user) VALUES
(1, 'dashboard', '-1', '概览', 'C', '/', 'dashboard/Overview', 'DashboardOutlined', 'sys:dashboard', 1, 1, '000000', 1),
(2, 'system', '-1', '系统管理', 'M', '/system', NULL, 'SettingOutlined', NULL, 2, 1, '000000', 1),
(3, 'user', 'system', '用户管理', 'C', '/system/users', 'system/UserManagement', 'TeamOutlined', 'sys:user:list', 1, 1, '000000', 1),
(4, 'dept', 'system', '部门管理', 'C', '/system/depts', 'system/DepartmentManagement', 'ApartmentOutlined', 'sys:dept:list', 2, 1, '000000', 1),
(5, 'post', 'system', '岗位管理', 'C', '/system/posts', 'system/PostManagement', 'IdcardOutlined', 'sys:post:list', 3, 1, '000000', 1),
(6, 'menu', 'system', '菜单管理', 'C', '/system/menus', 'system/MenuManagement', 'MenuOutlined', 'sys:menu:list', 4, 1, '000000', 1),
(7, 'role', 'system', '角色管理', 'C', '/system/roles', 'system/RoleManagement', 'SafetyCertificateOutlined', 'sys:role:list', 5, 1, '000000', 1),
(8, 'security', '-1', '安全审计', 'M', '/security', NULL, 'FileSearchOutlined', NULL, 3, 1, '000000', 1),
(9, 'login_log', 'security', '登录日志', 'C', '/security/login-log', 'security/LoginLog', 'FileTextOutlined', 'sys:loginlog:list', 1, 1, '000000', 1),
(10, 'oper_log', 'security', '操作日志', 'C', '/security/op-log', 'security/OperationLog', 'FileTextOutlined', 'sys:operlog:list', 2, 1, '000000', 1),
(11, 'online', 'security', '在线用户', 'C', '/security/online', 'security/OnlineUsers', 'TeamOutlined', 'sys:online:list', 3, 1, '000000', 1),
(12, 'monitor', '-1', '系统监控', 'M', '/monitor', NULL, 'DashboardOutlined', NULL, 4, 1, '000000', 1),
(13, 'server', 'monitor', '服务监控', 'C', '/monitor/server', 'monitor/ServiceMonitoring', 'DashboardOutlined', 'sys:server:list', 1, 1, '000000', 1),
(14, 'cache', 'monitor', '缓存监控', 'C', '/monitor/cache', 'monitor/CacheMonitoring', 'DatabaseOutlined', 'sys:cache:list', 2, 1, '000000', 1),
(15, 'datasource', 'monitor', '数据源监控', 'C', '/monitor/datasource', 'monitor/DataSourceMonitoring', 'DatabaseOutlined', 'sys:db:list', 3, 1, '000000', 1),
(16, 'cron', 'monitor', '任务调度', 'C', '/monitor/cron', 'monitor/CronJobs', 'ScheduleOutlined', 'sys:job:list', 4, 1, '000000', 1),
(17, 'infra', '-1', '基础设施', 'M', '/infrastructure', NULL, 'FolderOpenOutlined', NULL, 5, 1, '000000', 1),
(18, 'dict', 'infra', '字典管理', 'C', '/infrastructure/dict', 'infrastructure/Dictionary', 'SettingOutlined', 'sys:dict:list', 1, 1, '000000', 1),
(19, 'param', 'infra', '参数配置', 'C', '/infrastructure/params', 'infrastructure/Params', 'SettingOutlined', 'sys:config:list', 2, 1, '000000', 1),
(20, 'file', 'infra', '文件管理', 'C', '/infrastructure/files', 'infrastructure/FileCenter', 'FolderOpenOutlined', 'sys:file:list', 3, 1, '000000', 1),
(21, 'notice', 'infra', '通知公告', 'C', '/infrastructure/notices', 'infrastructure/Notices', 'NotificationOutlined', 'sys:notice:list', 4, 1, '000000', 1),
(22, 'devtools', '-1', '开发工具', 'M', '/devtools', NULL, 'CodeOutlined', NULL, 6, 1, '000000', 1),
(23, 'codegen', 'devtools', '代码生成', 'C', '/devtools/codegen', 'devtools/Codegen', 'CodeOutlined', 'dev:codegen:list', 1, 1, '000000', 1),
(24, 'api_doc', 'devtools', '接口文档', 'C', '/devtools/api-doc', 'devtools/ApiDoc', 'FileSearchOutlined', 'dev:api:list', 2, 1, '000000', 1),
(25, 'sql_terminal', 'devtools', 'SQL 终端', 'C', '/devtools/sql', 'devtools/SqlTerminal', 'DatabaseOutlined', 'dev:sql:list', 3, 1, '000000', 1),
(26, 'user_query', 'user', '用户查询', 'F', NULL, NULL, NULL, 'sys:user:query', 1, 1, '000000', 1),
(27, 'user_add', 'user', '用户新增', 'F', NULL, NULL, NULL, 'sys:user:add', 2, 1, '000000', 1),
(28, 'user_edit', 'user', '用户修改', 'F', NULL, NULL, NULL, 'sys:user:edit', 3, 1, '000000', 1),
(29, 'user_remove', 'user', '用户删除', 'F', NULL, NULL, NULL, 'sys:user:remove', 4, 1, '000000', 1);

-- ===============================
-- 6. 角色-菜单关联 (使用编码关联)
-- ===============================
INSERT INTO "zephyr_sys_role_menu" (role_code, menu_code) VALUES
('admin', 'dashboard'), ('admin', 'system'), ('admin', 'user'), ('admin', 'dept'), ('admin', 'post'), ('admin', 'menu'), ('admin', 'role'),
('admin', 'security'), ('admin', 'login_log'), ('admin', 'oper_log'), ('admin', 'online'),
('admin', 'monitor'), ('admin', 'server'), ('admin', 'cache'), ('admin', 'datasource'), ('admin', 'cron'),
('admin', 'infra'), ('admin', 'dict'), ('admin', 'param'), ('admin', 'file'), ('admin', 'notice'),
('admin', 'devtools'), ('admin', 'codegen'), ('admin', 'api_doc'), ('admin', 'sql_terminal'),
('admin', 'user_query'), ('admin', 'user_add'), ('admin', 'user_edit'), ('admin', 'user_remove'),
('common', 'dashboard'), ('common', 'system'), ('common', 'user'), ('common', 'user_query');