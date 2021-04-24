/*
 Navicat MySQL Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : 127.0.0.1:3306
 Source Schema         : web_api

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 23/04/2021 18:30:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL,
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `href` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int(11) NULL DEFAULT NULL,
  `target` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `state` tinyint(1) NULL DEFAULT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 36 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_menu
-- ----------------------------
INSERT INTO `system_menu` VALUES (1, 0, '常规管理', 'fa fa-address-book', '', 1, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (2, 0, '组件管理', 'fa fa-lemon-o', '', 2, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (3, 0, '其它管理', 'fa fa-slideshare', '', 3, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (5, 4, '主页一', 'fa fa-tachometer', 'page/welcome-1.html', 1, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (6, 4, '主页二', 'fa fa-tachometer', 'page/welcome-2.html', 2, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (7, 4, '主页三', 'fa fa-tachometer', 'page/welcome-3.html', 3, '_self', NULL, 1, '2021-04-23 09:54:59');
INSERT INTO `system_menu` VALUES (4, 1, '主页模板', 'fa fa-home', NULL, 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (8, 1, '菜单管理', 'fa fa-window-maximize', 'page/system/menu.html', 2, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (9, 1, '系统设置', 'fa fa-gears', 'page/setting.html', 3, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (10, 1, '表格示例', 'fa fa-file-text', 'page/table.html', 4, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (11, 1, '表单示例', 'fa fa-calendar', '', 5, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (12, 11, '普通表单', 'fa fa-list-alt', 'page/form.html', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (13, 11, '分步表单', 'fa fa-list-alt', 'page/form-step.html', 2, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (14, 1, '登录模板', 'fa fa-flag-o', '', 6, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (15, 14, '登录-1', 'fa fa-stumbleupon-circle', 'page/login-1.html', 1, '_blank', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (16, 14, '登录-2', 'fa fa-viacoin', 'page/login-2.html', 1, '_blank', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (17, 14, '登录-3', 'fa fa-tags', 'page/login-3.html', 1, '_blank', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (18, 1, '异常页面', 'fa fa-home', '', 7, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (19, 18, '404页面', 'fa fa-hourglass-end', 'page/404.html', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (20, 1, '其它界面', 'fa fa-snowflake-o', '', 8, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (21, 20, '按钮示例', 'fa fa-snowflake-o', 'page/button.html', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (22, 20, '弹出层', 'fa fa-shield', 'page/layer.html', 2, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (23, 2, '图标列表', 'fa fa-dot-circle-o', 'page/icon.html', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (24, 2, '图标选择', 'fa fa-adn', 'page/icon-picker.html', 2, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (25, 2, '颜色选择', 'fa fa-dashboard', 'page/color-select.html', 3, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (26, 2, '下拉选择', 'fa fa-angle-double-down', 'page/table-select.html', 4, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (27, 2, '文件上传', 'fa fa-arrow-up', 'page/upload.html', 5, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (28, 2, '富文本编辑器', 'fa fa-edit', 'page/editor.html', 6, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (29, 2, '省市县区选择器', 'fa fa-rocket', 'page/area.html', 7, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (30, 3, '多级菜单', 'fa fa-meetup', '', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (31, 3, '失效菜单', 'fa fa-superpowers', 'page/error.html', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (32, 30, '按钮1', 'fa fa-calendar', '', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (33, 32, '按钮2', 'fa fa-snowflake-o', '', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (34, 33, '按钮3', 'fa fa-snowflake-o', 'page/button.html?v=3', 1, '_self', NULL, 1, '2021-04-23 10:34:34');
INSERT INTO `system_menu` VALUES (35, 33, '按钮4', 'fa fa-calendar', 'page/form.html?v=1', 2, '_self', NULL, 1, '2021-04-23 10:34:34');

-- ----------------------------
-- Table structure for system_permission_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_permission_menu`;
CREATE TABLE `system_permission_menu`  (
  `permission_id` int(11) NOT NULL,
  `menu_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_permission_menu
-- ----------------------------
INSERT INTO `system_permission_menu` VALUES (1, '1,2,3,5,6,7,4,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35', NULL, NULL);

-- ----------------------------
-- Table structure for system_permissions
-- ----------------------------
DROP TABLE IF EXISTS `system_permissions`;
CREATE TABLE `system_permissions`  (
  `permission_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_permissions
-- ----------------------------
INSERT INTO `system_permissions` VALUES (1, '所有权限', '', NULL, NULL, '2021-04-23 14:58:27', '2021-04-23 14:58:29');

-- ----------------------------
-- Table structure for system_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `system_role_permission`;
CREATE TABLE `system_role_permission`  (
  `role_id` int(11) NOT NULL,
  `permission_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_role_permission
-- ----------------------------
INSERT INTO `system_role_permission` VALUES (1, '1', NULL, NULL);

-- ----------------------------
-- Table structure for system_roles
-- ----------------------------
DROP TABLE IF EXISTS `system_roles`;
CREATE TABLE `system_roles`  (
  `role_id` int(11) NOT NULL,
  `role_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_roles
-- ----------------------------
INSERT INTO `system_roles` VALUES (1, '超级管理员', '2021-04-23 14:58:14', '2021-04-23 14:58:16');

-- ----------------------------
-- Table structure for system_user_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_user_menu`;
CREATE TABLE `system_user_menu`  (
  `user_id` int(11) NOT NULL,
  `menu_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_user_menu
-- ----------------------------

-- ----------------------------
-- Table structure for system_user_permission
-- ----------------------------
DROP TABLE IF EXISTS `system_user_permission`;
CREATE TABLE `system_user_permission`  (
  `user_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Records of system_user_permission
-- ----------------------------

-- ----------------------------
-- Table structure for system_users
-- ----------------------------
DROP TABLE IF EXISTS `system_users`;
CREATE TABLE `system_users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass_word` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_id` int(11) NULL DEFAULT NULL,
  `avatar` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_users
-- ----------------------------
INSERT INTO `system_users` VALUES (1, 'admin', '$2y$10$rPK.GV05apkfe9nmbjfqtOevKQrdGEm2v2fUXIFmILqZBSxVj8cP.', 'Administrator', 1, NULL, '3TUJNzSS7AtOYRCQy7NuEW2u6lar8PXXcdTVdiZqHhgUFkkk2kG0Kk4EDTMJ', '2021-03-24 01:06:47', '2021-03-24 01:06:47');

SET FOREIGN_KEY_CHECKS = 1;
