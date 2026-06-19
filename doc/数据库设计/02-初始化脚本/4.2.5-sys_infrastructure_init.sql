-- ----------------------------
-- 基础设施模块 (Infrastructure) 初始化脚本
-- 涵盖常用的系统字典与参数默认值
-- ----------------------------




-- ----------------------------
-- 1. 字典类型数据
-- ----------------------------
INSERT INTO "zephyr_sys_dict_type" ("id", "dict_name", "dict_type", "status", "remark", "tenant_code", "create_user") VALUES 
(1, '用户性别', 'sys_user_sex', 0, '用户性别列表', '000000', 1),
(2, '系统状态', 'sys_common_status', 0, '通用状态列表', '000000', 1),
(3, '通知类型', 'sys_notice_type', 0, '系统通知类型', '000000', 1),
(4, '存储类型', 'sys_store_type', 0, '文件存储驱动类型', '000000', 1);

-- ----------------------------
-- 2. 字典键值数据
-- ----------------------------
-- 性别 (sys_user_sex)
INSERT INTO "zephyr_sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "list_class", "is_default", "status", "tenant_code", "create_user") VALUES 
(101, 1, '男', '1', 'sys_user_sex', 'primary', 1, 0, '000000', 1),
(102, 2, '女', '2', 'sys_user_sex', 'danger', 0, 0, '000000', 1),
(103, 3, '未知', '0', 'sys_user_sex', 'info', 0, 0, '000000', 1);

-- 通用状态 (sys_common_status)
INSERT INTO "zephyr_sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "list_class", "is_default", "status", "tenant_code", "create_user") VALUES 
(201, 1, '正常', '0', 'sys_common_status', 'success', 1, 0, '000000', 1),
(202, 2, '停用', '1', 'sys_common_status', 'danger', 0, 0, '000000', 1);

-- 通知类型 (sys_notice_type)
INSERT INTO "zephyr_sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "list_class", "is_default", "status", "tenant_code", "create_user") VALUES 
(301, 1, '通知', '1', 'sys_notice_type', 'info', 1, 0, '000000', 1),
(302, 2, '公告', '2', 'sys_notice_type', 'warning', 0, 0, '000000', 1);

-- 存储类型 (sys_store_type)
INSERT INTO "zephyr_sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "list_class", "is_default", "status", "tenant_code", "create_user") VALUES 
(401, 1, '本地存储', 'local', 'sys_store_type', 'primary', 1, 0, '000000', 1),
(402, 2, '阿里云OSS', 'oss', 'sys_store_type', 'success', 0, 0, '000000', 1),
(403, 3, 'Minio', 'minio', 'sys_store_type', 'warning', 0, 0, '000000', 1);

-- ----------------------------
-- 3. 系统参数预置
-- ----------------------------
INSERT INTO "zephyr_sys_config" ("id", "config_name", "config_key", "config_value", "config_type", "remark", "tenant_code", "create_user") VALUES 
(1, '用户初始化密码', 'sys.user.initPassword', '123456', 1, '用户管理界面新增用户时的默认密码', '000000', 1),
(2, '系统侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 1, '侧边栏主题类型：theme-dark/theme-light', '000000', 1),
(3, '文件上传限制', 'sys.file.uploadLimit', '10', 1, '单个文件上传最大限制（MB）', '000000', 1);

-- ----------------------------
-- 4. 示例通知公告
-- ----------------------------
INSERT INTO "zephyr_sys_notice" ("id", "notice_title", "notice_type", "notice_content", "status", "tenant_code", "create_user") VALUES 
(1, 'Zephyr 系统发版公告', 1, '<p>欢迎使用 Zephyr 管理后台！当前版本已集成核心监控与安全设计。</p>', 1, '000000', 1);


