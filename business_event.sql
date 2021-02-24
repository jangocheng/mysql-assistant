/*
 Navicat Premium Data Transfer

 Source Server         : my-rds
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : rm-2zemxuvee9kii2b55so.mysql.rds.aliyuncs.com:3306
 Source Schema         : business_event

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 24/02/2021 09:37:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `administrator_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '缩略图',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`administrator_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='后台管理人员';

-- ----------------------------
-- Records of administrator
-- ----------------------------
BEGIN;
INSERT INTO `administrator` VALUES (1, 'admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', '主管理帐号', NULL, '2019-05-23 12:09:26', '2020-07-21 22:47:03', 0, '2020-07-20 14:53:30');
INSERT INTO `administrator` VALUES (3, 'owen', '8cb2237d0679ca88db6464eac60da96345513964', 'show me', NULL, '2019-05-23 12:09:26', '2020-07-22 10:21:48', 0, '2020-07-20 14:53:18');
COMMIT;

-- ----------------------------
-- Table structure for ddd_event
-- ----------------------------
DROP TABLE IF EXISTS `ddd_event`;
CREATE TABLE `ddd_event` (
  `ddd_event_id` int(11) NOT NULL AUTO_INCREMENT,
  `event_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'mysql',
  `event_tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '事件标签',
  `event_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '事件名称',
  `stream_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '事件包含的mysql操作',
  `event_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `event_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '事件外部文档',
  `comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '事件说明',
  `is_deleted` int(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ddd_event_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for ddd_event_stream
-- ----------------------------
DROP TABLE IF EXISTS `ddd_event_stream`;
CREATE TABLE `ddd_event_stream` (
  `ddd_event_stream_id` int(11) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `transaction_tag` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `event_type` int(11) DEFAULT '-100' COMMENT '1insert, 2update, 3delete',
  `columns` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `update_columns` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '更新的字段',
  `update_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '更新字段的值',
  `ignore_column_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '忽略的字段值',
  `comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ddd_event_stream_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=214748 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `menu_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '{"name":"文档ID","desc":"哈哈哈哈哈哈","type":"password"}',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '{"name":"上级ID","desc":"","type":"select", "options":{"callback":"getMenuTree"}}',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序（同级有效）',
  `hide` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '{"name":"是否隐藏","options":{"1":"否","2": "是"}}',
  `pathname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '路由',
  `iconfont` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '{"name":"图标"}',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=241 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (200, '菜单列表', 0, 0, 0, '/admin/menu/list.html', 'fa-bars', '2020-02-16 09:14:38', '2021-02-09 10:35:08', 0, '2020-07-21 13:09:05');
INSERT INTO `menu` VALUES (201, '管理员列表', 0, 0, 0, '/admin/admins/list.html', 'fa-user', '2020-02-16 09:14:38', '2021-02-11 07:43:21', 1, '2021-02-11 07:43:21');
INSERT INTO `menu` VALUES (203, '资源列表', 0, 0, 0, '/admin/resource/list.html', 'fa-tag', '2020-02-16 09:14:38', '2021-02-11 07:43:14', 1, '2021-02-11 07:43:14');
INSERT INTO `menu` VALUES (205, '权限管理', 0, 0, 0, '', 'fa-share', '2020-02-16 09:14:38', '2021-02-11 07:42:55', 1, '2021-02-11 07:42:55');
INSERT INTO `menu` VALUES (206, '权限节点列表', 205, 0, 0, '/admin/permissions/list.html', 'fa-tag', '2020-02-16 09:14:38', '2020-07-22 13:48:33', 0, NULL);
INSERT INTO `menu` VALUES (207, '角色列表', 205, 0, 0, '/admin/roles/list.html', 'fa-tag', '2020-02-16 09:14:38', '2020-07-22 13:48:36', 0, '2020-07-20 22:42:33');
INSERT INTO `menu` VALUES (231, '业务事件数据模型', 0, 0, 0, '', '', '2020-11-12 18:11:19', '2020-11-15 20:38:55', 0, NULL);
INSERT INTO `menu` VALUES (232, '事件列表', 231, 0, 0, '/admin/event/list.html', '', '2020-11-12 18:11:52', '2020-11-12 18:14:26', 0, NULL);
INSERT INTO `menu` VALUES (233, 'binlog数据流', 231, 0, 0, '/admin/event/stream_list.html', '', '2020-11-12 18:26:20', '2020-11-15 20:41:49', 0, NULL);
INSERT INTO `menu` VALUES (234, '业务状态管理', 0, 100, 0, '', '', '2021-02-09 10:36:26', '2021-02-09 10:37:00', 0, NULL);
INSERT INTO `menu` VALUES (235, '业务状态定义', 234, 0, 0, '/admin/state/list.html', '', '2021-02-09 10:38:17', '2021-02-09 11:24:08', 0, NULL);
INSERT INTO `menu` VALUES (236, '数据变更统计', 0, 0, 0, '', '', '2021-02-09 11:22:21', '2021-02-09 11:22:21', 0, NULL);
INSERT INTO `menu` VALUES (237, '统计规则', 236, 0, 0, '/admin/statistics/rule_list.html', '', '2021-02-09 11:22:47', '2021-02-18 23:03:58', 0, NULL);
INSERT INTO `menu` VALUES (238, '异常状态变更', 234, 0, 0, '/admin/state/abnormal_list.html', '', '2021-02-18 20:15:36', '2021-02-18 20:15:36', 0, NULL);
INSERT INTO `menu` VALUES (239, '每日统计列表', 236, 0, 0, '/admin/statistics/day_list.html', '', '2021-02-18 22:05:32', '2021-02-19 21:12:06', 0, '2021-02-19 21:11:32');
INSERT INTO `menu` VALUES (240, 'aaa', 0, 0, 0, '', '', '0000-00-00 00:00:00', '2021-02-19 21:17:59', 1, '2021-02-19 21:17:59');
COMMIT;

-- ----------------------------
-- Table structure for rbac_permission
-- ----------------------------
DROP TABLE IF EXISTS `rbac_permission`;
CREATE TABLE `rbac_permission` (
  `permission_id` int(11) NOT NULL AUTO_INCREMENT,
  `method` char(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `source` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `title` char(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`permission_id`) USING BTREE,
  KEY `Title` (`title`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role`;
CREATE TABLE `rbac_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE,
  KEY `Title` (`title`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of rbac_role
-- ----------------------------
BEGIN;
INSERT INTO `rbac_role` VALUES (1, 'root', '超级管理员', 0, NULL);
INSERT INTO `rbac_role` VALUES (2, '基础权限', '没此基础权限的管理员，不能以常规方式登录系统', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for rbac_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_permission`;
CREATE TABLE `rbac_role_permission` (
  `role_permissions_id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `assignment_date` datetime NOT NULL,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`role_permissions_id`) USING BTREE,
  UNIQUE KEY `role_permission` (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for rbac_user_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_user_role`;
CREATE TABLE `rbac_user_role` (
  `user_role_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `assignment_date` datetime NOT NULL,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_role_id`) USING BTREE,
  UNIQUE KEY `user_role` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of rbac_user_role
-- ----------------------------
BEGIN;
INSERT INTO `rbac_user_role` VALUES (19, 3, 1, '2019-05-30 23:35:24', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (20, 2, 4, '2020-02-16 12:52:12', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (21, 2, 3, '2020-02-16 12:52:12', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (22, 2, 1, '2020-02-16 12:52:12', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (30, 1, 1, '0000-00-00 00:00:00', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (31, 1, 2, '0000-00-00 00:00:00', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (32, 1, 4, '0000-00-00 00:00:00', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (33, 1, 5, '0000-00-00 00:00:00', 0, NULL);
INSERT INTO `rbac_user_role` VALUES (34, 1, 6, '0000-00-00 00:00:00', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `resource_id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_group` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '分组',
  `resource_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `identity` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
BEGIN;
INSERT INTO `resource` VALUES (1, 'admin', '资源', 'Sourcessss', 0, NULL);
INSERT INTO `resource` VALUES (2, 'admin', 'rbac-权限', 'Permissions', 0, NULL);
INSERT INTO `resource` VALUES (3, 'admin', '角色', 'Roles', 0, NULL);
INSERT INTO `resource` VALUES (4, 'admin', '侧边栏', 'Sidebar', 0, NULL);
INSERT INTO `resource` VALUES (5, 'qqss', 'w', 'eess', 1, '2020-07-21 00:16:56');
COMMIT;

-- ----------------------------
-- Table structure for state
-- ----------------------------
DROP TABLE IF EXISTS `state`;
CREATE TABLE `state` (
  `state_id` int(11) NOT NULL AUTO_INCREMENT,
  `state_class_id` int(11) DEFAULT NULL,
  `state_value` varchar(255) DEFAULT NULL,
  `state_value_desc` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`state_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of state
-- ----------------------------
BEGIN;
INSERT INTO `state` VALUES (1, 1, '1', '待支付', '2021-02-09 11:34:37', NULL, 0, NULL);
INSERT INTO `state` VALUES (2, 1, '2', '部分支付', '2021-02-09 11:35:01', '2021-02-09 11:35:03', 0, NULL);
INSERT INTO `state` VALUES (3, 1, '3', '支付完成', '2021-02-09 11:35:12', '2021-02-09 11:35:41', 0, NULL);
INSERT INTO `state` VALUES (4, 1, '4', '异步回调失败', '2021-02-09 11:35:40', NULL, 0, NULL);
INSERT INTO `state` VALUES (5, 1, '5', '异步回调成功，待发货', '2021-02-09 11:36:23', NULL, 0, NULL);
INSERT INTO `state` VALUES (6, 1, '6', '已发货', '2021-02-09 11:36:43', NULL, 0, NULL);
INSERT INTO `state` VALUES (7, 1, '7', '已签收', '2021-02-09 11:37:07', NULL, 0, NULL);
INSERT INTO `state` VALUES (8, 1, '8', '退款中-1', '2021-02-09 11:37:22', '2021-02-10 22:53:59', 0, '2021-02-10 21:00:34');
INSERT INTO `state` VALUES (11, 3, '0', '默认', '2021-02-18 17:10:33', '2021-02-19 15:35:01', 1, '2021-02-19 15:35:01');
INSERT INTO `state` VALUES (12, 3, '1', '活动开始', '2021-02-18 17:10:49', NULL, 0, NULL);
INSERT INTO `state` VALUES (13, 3, '2', '活动进行中', '2021-02-18 17:11:01', '2021-02-19 17:42:57', 0, NULL);
INSERT INTO `state` VALUES (14, 3, '3', '活动结束', '2021-02-18 17:11:11', '2021-02-19 17:28:44', 1, '2021-02-19 17:28:44');
COMMIT;

-- ----------------------------
-- Table structure for state_abnormal
-- ----------------------------
DROP TABLE IF EXISTS `state_abnormal`;
CREATE TABLE `state_abnormal` (
  `state_abnormal_id` int(11) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(255) DEFAULT '',
  `table_name` varchar(255) DEFAULT '',
  `field_name` varchar(255) DEFAULT '',
  `event_type` int(11) DEFAULT '0',
  `state_from` varchar(255) DEFAULT NULL,
  `state_to` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`state_abnormal_id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8 COMMENT='这个表的记录，考拉两种方式\n1， mysql-binlog解析时 就校验其是否正确\n\n2， 把mysql-binlog解析完成的的结果存表， 统计脚本对结果进行分析。\n\n方案1更合理，  方案2可以用大数据的工具，进行分析。   ';

-- ----------------------------
-- Records of state_abnormal
-- ----------------------------
BEGIN;
INSERT INTO `state_abnormal` VALUES (1, 'codeper', 'activity', 'status', 0, '2', '22', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for state_class
-- ----------------------------
DROP TABLE IF EXISTS `state_class`;
CREATE TABLE `state_class` (
  `state_class_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `state_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `db_name` varchar(255) DEFAULT '',
  `table_name` varchar(255) DEFAULT '',
  `field_name` varchar(255) DEFAULT '',
  `state_describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `status` tinyint(4) DEFAULT '1',
  `is_deleted` tinyint(4) DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`state_class_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of state_class
-- ----------------------------
BEGIN;
INSERT INTO `state_class` VALUES (1, '订单状态', 'test', 'order', 'order_status', '订单流程说明adada', 1, 0, '2021-02-10 23:05:50', '2021-02-10 23:05:50', NULL);
INSERT INTO `state_class` VALUES (2, 'test', 'sss', 'bbb', 'dd', '', 1, 1, '2021-02-10 23:03:12', '2021-02-10 23:03:12', '2021-02-10 23:03:12');
INSERT INTO `state_class` VALUES (3, '活动状态', 'codeper', 'activity', 'status', '123123123', 1, 0, '2021-02-19 17:39:10', '2021-02-19 17:39:10', NULL);
INSERT INTO `state_class` VALUES (4, '测试', 'a', 'b', 'c', '', 1, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for state_direction
-- ----------------------------
DROP TABLE IF EXISTS `state_direction`;
CREATE TABLE `state_direction` (
  `state_direction_id` int(11) NOT NULL AUTO_INCREMENT,
  `state_class_id` int(11) DEFAULT '0',
  `state_from` varchar(255) DEFAULT '0',
  `state_to` varchar(255) DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`state_direction_id`),
  UNIQUE KEY `dir_idx` (`state_class_id`,`state_from`,`state_to`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of state_direction
-- ----------------------------
BEGIN;
INSERT INTO `state_direction` VALUES (1, 1, '1', '2', '2021-02-10 05:40:00', '2021-02-10 20:41:09', 0, '2021-02-10 05:40:00');
INSERT INTO `state_direction` VALUES (2, 1, '1', '3', '2021-02-10 05:40:37', '2021-02-10 20:41:11', 0, '2021-02-10 05:40:37');
INSERT INTO `state_direction` VALUES (3, 1, '2', '3', '2021-02-10 05:40:42', '2021-02-10 20:41:12', 0, '2021-02-10 05:40:42');
INSERT INTO `state_direction` VALUES (5, 1, '3', '4', '2021-02-10 20:42:33', '2021-02-10 20:42:33', 0, '2021-02-10 20:42:33');
INSERT INTO `state_direction` VALUES (6, 1, '3', '5', '2021-02-10 20:42:43', '2021-02-10 20:42:43', 0, '2021-02-10 20:42:43');
INSERT INTO `state_direction` VALUES (7, 1, '5', '6', '2021-02-10 20:42:54', '2021-02-10 20:46:52', 0, '2021-02-10 20:42:54');
INSERT INTO `state_direction` VALUES (8, 1, '6', '7', '2021-02-10 20:43:22', '2021-02-10 20:47:05', 0, '2021-02-10 20:43:22');
INSERT INTO `state_direction` VALUES (9, 1, '5', '8', '2021-02-10 20:46:20', '2021-02-10 20:47:35', 0, '2021-02-10 20:46:20');
INSERT INTO `state_direction` VALUES (10, 1, '7', '8', '2021-02-10 20:47:44', '2021-02-10 20:57:21', 0, '2021-02-10 20:56:19');
INSERT INTO `state_direction` VALUES (11, 3, '0', '1', '2021-02-18 17:14:05', '2021-02-19 15:34:45', 1, '2021-02-19 15:34:45');
INSERT INTO `state_direction` VALUES (12, 3, '1', '2', '2021-02-18 17:14:22', '2021-02-18 17:14:22', 0, '2021-02-18 17:14:22');
INSERT INTO `state_direction` VALUES (13, 3, '2', '3', '2021-02-18 17:16:00', '2021-02-18 17:16:00', 0, '2021-02-18 17:16:00');
INSERT INTO `state_direction` VALUES (14, 3, '1', '3', '2021-02-18 17:16:22', '2021-02-18 17:16:22', 0, '2021-02-18 17:16:22');
COMMIT;

-- ----------------------------
-- Table structure for statistics_day
-- ----------------------------
DROP TABLE IF EXISTS `statistics_day`;
CREATE TABLE `statistics_day` (
  `statistics_day_id` int(11) NOT NULL AUTO_INCREMENT,
  `statistics_rule_id` int(11) DEFAULT '0',
  `statistics_day` date DEFAULT NULL,
  `db_name` varchar(255) DEFAULT '',
  `table_name` varchar(255) DEFAULT '',
  `field_name` varchar(255) DEFAULT '',
  `insert_times` int(11) unsigned DEFAULT '0',
  `update_times` int(11) unsigned DEFAULT '0',
  `delete_times` int(11) unsigned DEFAULT '0',
  `is_deleted` tinyint(4) DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`statistics_day_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of statistics_day
-- ----------------------------
BEGIN;
INSERT INTO `statistics_day` VALUES (15, 2, '2021-02-19', 'codeper', 'activity', 'status', 0, 10, 0, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `statistics_day` VALUES (16, 1, '2021-02-19', 'codeper', 'activity', '', 3, 14, 4, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for statistics_rule
-- ----------------------------
DROP TABLE IF EXISTS `statistics_rule`;
CREATE TABLE `statistics_rule` (
  `statistics_rule_id` int(11) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(255) DEFAULT '',
  `table_name` varchar(255) DEFAULT '',
  `field_name` varchar(255) DEFAULT '',
  `is_deleted` tinyint(4) DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`statistics_rule_id`),
  UNIQUE KEY `dbf_idx` (`db_name`,`table_name`,`field_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of statistics_rule
-- ----------------------------
BEGIN;
INSERT INTO `statistics_rule` VALUES (1, 'codeper', 'activity', '', 0, '2021-02-18 23:11:19', '2021-02-18 23:11:19', NULL);
INSERT INTO `statistics_rule` VALUES (2, 'codeper', 'activity', 'status', 0, '2021-02-18 23:11:35', '2021-02-19 15:02:04', '2021-02-19 15:01:54');
INSERT INTO `statistics_rule` VALUES (3, 'codeper', 'haha', '', 1, '0000-00-00 00:00:00', '2021-02-19 21:25:13', '2021-02-19 21:25:13');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
