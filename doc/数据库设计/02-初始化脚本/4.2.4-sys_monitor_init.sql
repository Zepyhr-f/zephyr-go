/* -- ----------------------------
-- 1. 安全审计顶级菜单 (已在 RBAC_init.sql 中初始化，避免重复)
-- ----------------------------
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (200, 'audit', '-1', '安全审计', 'M', '/audit', NULL, NULL, 'safetyCertificate', 5, 1, '000000', 1, NOW());

-- 2.1 登录日志
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (201, 'login_log', 'audit', '登录日志', 'C', 'login', 'system/login-log', 'sys:loginlog:list', 'login', 1, 1, '000000', 1, NOW());

-- 2.2 操作日志
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (202, 'oper_log', 'audit', '操作日志', 'C', 'oper', 'system/oper-log', 'sys:operlog:list', 'file-text', 2, 1, '000000', 1, NOW());

-- 2.3 在线用户
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (203, 'online_user', 'audit', '在线用户', 'C', 'online', 'system/online', 'sys:online:list', 'user', 3, 1, '000000', 1, NOW());
*/

-- ----------------------------
-- 2. 系统监控顶级菜单
-- ----------------------------
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (300, 'monitor', '-1', '系统监控', 'M', '/monitor', NULL, NULL, 'dashboard', 6, 1, '000000', 1, NOW());

-- 3.1 服务监控
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (301, 'server_monitor', 'monitor', '服务监控', 'C', 'server', 'monitor/server', 'monitor:server:list', 'chart', 1, 1, '000000', 1, NOW());

-- 3.2 缓存监控
INSERT INTO zephyr_sys_menu (id, code, parent_code, menu_name, menu_type, path, component, perms, icon, order_num, status, tenant_code, create_user, create_time)
VALUES (302, 'cache_monitor', 'monitor', '缓存监控', 'C', 'cache', 'monitor/cache', 'monitor:cache:list', 'database', 2, 1, '000000', 1, NOW());

-- ----------------------------
-- 3. 系统内置管理员授权 (Role Code='admin')
-- ----------------------------
INSERT INTO zephyr_sys_role_menu (role_code, menu_code) VALUES 
('admin', 'monitor'), ('admin', 'server_monitor'), ('admin', 'cache_monitor');
