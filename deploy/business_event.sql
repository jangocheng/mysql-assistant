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

 Date: 01/01/2021 17:46:54
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
  `status` tinyint(1) DEFAULT NULL,
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
INSERT INTO `administrator` VALUES (1, 'admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', '主管理帐号', NULL, 33, '2019-05-23 12:09:26', '2020-07-21 22:47:03', 0, '2020-07-20 14:53:30');
INSERT INTO `administrator` VALUES (3, 'owen', '8cb2237d0679ca88db6464eac60da96345513964', 'show me', NULL, 1, '2019-05-23 12:09:26', '2020-07-22 10:21:48', 0, '2020-07-20 14:53:18');
INSERT INTO `administrator` VALUES (6, 'test', NULL, 'aaaaa', NULL, 1, NULL, '2020-09-20 10:12:18', 1, '2020-09-20 02:12:18');
INSERT INTO `administrator` VALUES (7, 'www', NULL, 'wwwwqqaass', NULL, 127, '2020-07-20 22:23:13', '2020-07-21 21:16:37', 0, '2020-07-20 14:52:45');
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for ddd_event_stream
-- ----------------------------
DROP TABLE IF EXISTS `ddd_event_stream`;
CREATE TABLE `ddd_event_stream` (
  `ddd_event_stream_id` int(11) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `transaction_tag` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `event_type` int(11) DEFAULT '-100',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for ddd_event_stream_copy1
-- ----------------------------
DROP TABLE IF EXISTS `ddd_event_stream_copy1`;
CREATE TABLE `ddd_event_stream_copy1` (
  `ddd_event_stream_id` int(11) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `transaction_tag` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `event_type` int(11) DEFAULT '-100',
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
) ENGINE=InnoDB AUTO_INCREMENT=175357 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of ddd_event_stream_copy1
-- ----------------------------
BEGIN;
INSERT INTO `ddd_event_stream_copy1` VALUES (8719, 'eas_test', 'class_position_details_37', '5fb32c11afd72-1605577745', 1, 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"16464150\",\"class_id\":\"10010661\",\"class_position_id\":\"230004\",\"degree_no\":\"7\",\"serial_number\":\"14\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-16 15:49:07\",\"updated_at\":\"2020-11-16 15:49:07\",\"sys_update_dc\":\"2020-11-16 15:49:07\"}', NULL, NULL, 1, '2020-12-14 19:41:20', '2020-11-17 09:49:09', '2020-12-14 19:41:19');
INSERT INTO `ddd_event_stream_copy1` VALUES (8720, 'eas_test', 'class_position_details_37', '5fb32c11afd72-1605577745', 1, 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"16464151\",\"class_id\":\"10010661\",\"class_position_id\":\"230004\",\"degree_no\":\"7\",\"serial_number\":\"15\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-16 15:49:07\",\"updated_at\":\"2020-11-16 15:49:07\",\"sys_update_dc\":\"2020-11-16 15:49:07\"}', NULL, NULL, 1, '2020-12-14 19:41:58', '2020-11-17 09:49:09', '2020-12-14 19:41:57');
INSERT INTO `ddd_event_stream_copy1` VALUES (8721, 'eas_test', 'class_position_details_37', '', 1, 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"16464260\",\"class_id\":\"10010661\",\"class_position_id\":\"230009\",\"degree_no\":\"12\",\"serial_number\":\"4\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-16 15:49:07\",\"updated_at\":\"2020-11-16 15:49:07\",\"sys_update_dc\":\"2020-11-16 15:49:07\"}', NULL, NULL, 0, NULL, '2020-11-17 14:31:21', '2020-11-17 14:31:21');
INSERT INTO `ddd_event_stream_copy1` VALUES (8722, 'eas_test', 'class_position_details_37', '', 1, 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"16464261\",\"class_id\":\"10010661\",\"class_position_id\":\"230009\",\"degree_no\":\"12\",\"serial_number\":\"5\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-16 15:49:07\",\"updated_at\":\"2020-11-16 15:49:07\",\"sys_update_dc\":\"2020-11-16 15:49:07\"}', NULL, NULL, 0, NULL, '2020-11-17 14:31:21', '2020-11-17 14:31:21');
INSERT INTO `ddd_event_stream_copy1` VALUES (8723, 'eas_test', 'class_position_details_37', '', 1, 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,class_position_id,degree_no,serial_number,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"16464262\",\"class_id\":\"10010661\",\"class_position_id\":\"230009\",\"degree_no\":\"12\",\"serial_number\":\"6\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-16 15:49:07\",\"updated_at\":\"2020-11-16 15:49:07\",\"sys_update_dc\":\"2020-11-16 15:49:07\"}', NULL, NULL, 0, NULL, '2020-11-17 14:31:21', '2020-11-17 14:31:21');
INSERT INTO `ddd_event_stream_copy1` VALUES (30051, 'eas_test', 'classes', '5fb370d02c2fe-1605595344', 1, 'id,uid,campus_id,teaching_center_id,product_id,classroom_id,classroom_ids,time_template_id,name,dict_class_status_id,plan_min_class_number,plan_max_class_number,insert_class,promoted_status,coursed,sold_degrees,class_teacher_master_name,guidance_teacher,start_sale_time,start_attend_class_time,end_class_time,inside_sell_time,external_sell_time,is_have_promoted,is_original,is_short_term,is_online,is_demo,is_make_up,is_experiment,is_demo_class,demo_class_category_id,set_online_status,old_class_id,half_start_date,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,uid,campus_id,teaching_center_id,product_id,classroom_id,classroom_ids,time_template_id,name,dict_class_status_id,plan_min_class_number,plan_max_class_number,insert_class,promoted_status,coursed,sold_degrees,class_teacher_master_name,guidance_teacher,start_sale_time,start_attend_class_time,end_class_time,inside_sell_time,external_sell_time,is_have_promoted,is_original,is_short_term,is_online,is_demo,is_make_up,is_experiment,is_demo_class,demo_class_category_id,set_online_status,old_class_id,half_start_date,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"10010678\",\"uid\":\"110\",\"campus_id\":\"37\",\"teaching_center_id\":\"13\",\"product_id\":\"10001\",\"classroom_id\":\"0\",\"classroom_ids\":\"4344\",\"time_template_id\":\"0\",\"name\":\"D-EK1-BJBY-201117-1\",\"dict_class_status_id\":\"4\",\"plan_min_class_number\":\"15\",\"plan_max_class_number\":\"15\",\"insert_class\":\"0\",\"promoted_status\":\"0\",\"coursed\":\"0\",\"sold_degrees\":\"0\",\"class_teacher_master_name\":\"\",\"guidance_teacher\":\"\",\"start_sale_time\":\"2020-11-17 11:59:27\",\"start_attend_class_time\":\"2020-11-18 20:00:00\",\"end_class_time\":\"2020-11-18 20:55:00\",\"inside_sell_time\":\"\",\"external_sell_time\":\"\",\"is_have_promoted\":\"0\",\"is_original\":\"0\",\"is_short_term\":\"0\",\"is_online\":\"0\",\"is_demo\":\"0\",\"is_make_up\":\"0\",\"is_experiment\":\"0\",\"is_demo_class\":\"1\",\"demo_class_category_id\":\"2\",\"set_online_status\":\"0\",\"old_class_id\":\"\",\"half_start_date\":\"\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-17 11:59:27\",\"updated_at\":\"2020-11-17 11:59:27\",\"sys_update_dc\":\"2020-11-17 11:59:27\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:24', '2020-11-17 14:42:24');
INSERT INTO `ddd_event_stream_copy1` VALUES (30052, 'eas_test', 'class_teachers', '5fb370d02c2fe-1605595344', 1, 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"89629\",\"class_id\":\"10010678\",\"teacher_id\":\"2816\",\"dict_teacher_duty_id\":\"1\",\"dict_teacher_attr_id\":\"1\",\"teacher_sort\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-17 11:59:27\",\"updated_at\":\"2020-11-17 11:59:27\",\"sys_update_dc\":\"2020-11-17 11:59:27\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:24', '2020-11-17 14:42:24');
INSERT INTO `ddd_event_stream_copy1` VALUES (30053, 'contra_test_rw', 'contract_divide_detail_37', '5fb370de41e7c-1605595358', 2, 'id,student_id,student_no,contract_id,contract_detail_id,contract_detail_no,class_id,position_id,position_detail_id,degree_no,course_serial_id,status,is_history,created_at,updated_at,deleted_at,sys_update_dc,is_consume', 'status,updated_at,sys_update_dc', '{\"status\":\"3\",\"updated_at\":\"2020-11-17 13:47:39\",\"sys_update_dc\":\"2020-11-17 13:47:39\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:38', '2020-11-17 14:42:38');
INSERT INTO `ddd_event_stream_copy1` VALUES (30156, 'contra_test_rw', 'contract_detail_73', '5fb370de41e7c-1605595358', 2, 'id,contract_id,exec_status,standard_price,real_price,sign_campus_id,exec_campus_id,refund_contract_id,consume_id,contract_serial_id,class_serial_id,sign_entity,pay_entity,is_history,created_at,updated_at,deleted_at,sys_update_dc,change_tag', 'exec_status,updated_at,sys_update_dc', '{\"exec_status\":\"1\",\"updated_at\":\"2020-11-17 13:47:39\",\"sys_update_dc\":\"2020-11-17 13:47:39\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:40', '2020-11-17 14:42:40');
INSERT INTO `ddd_event_stream_copy1` VALUES (30157, 'contra_test_rw', 'contract', '5fb370de41e7c-1605595358', 2, 'id,type,status,parent_id,sign_time,student_id,student_no,campus_area_id,campus_id,product_id,product_name,class_num,standard_amount,real_amount,total_amount,finance_id,tax_id,related_id,not_divide_num,is_history,begin_serial_id,end_serial_id,class_serial_ids,exec_campus_id,parent_phone,pay_class_num,pay_serial_ids,step_serial_ids,end_time,encrypt,origin_campus_area_id,origin_campus_id,created_at,updated_at,deleted_at,sys_update_dc', 'status,not_divide_num,updated_at,sys_update_dc', '{\"status\":\"2\",\"not_divide_num\":\"80\",\"updated_at\":\"2020-11-17 13:47:39\",\"sys_update_dc\":\"2020-11-17 13:47:39\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:40', '2020-11-17 14:42:40');
INSERT INTO `ddd_event_stream_copy1` VALUES (30158, 'contra_test_rw', 'contract', '5fb370de41e7c-1605595358', 2, 'id,type,status,parent_id,sign_time,student_id,student_no,campus_area_id,campus_id,product_id,product_name,class_num,standard_amount,real_amount,total_amount,finance_id,tax_id,related_id,not_divide_num,is_history,begin_serial_id,end_serial_id,class_serial_ids,exec_campus_id,parent_phone,pay_class_num,pay_serial_ids,step_serial_ids,end_time,encrypt,origin_campus_area_id,origin_campus_id,created_at,updated_at,deleted_at,sys_update_dc', 'status,not_divide_num,updated_at,sys_update_dc', '{\"status\":\"2\",\"not_divide_num\":\"24\",\"updated_at\":\"2020-11-17 13:47:39\",\"sys_update_dc\":\"2020-11-17 13:47:39\"}', NULL, NULL, 0, NULL, '2020-11-17 14:42:41', '2020-11-17 14:42:41');
INSERT INTO `ddd_event_stream_copy1` VALUES (167858, 'eas_test', 'class_stats', '5fb78c5ba752d-1605864539', 1, 'class_id,continual_surplus_num,surplus_num,in_class_num,created_at,updated_at', 'class_id,continual_surplus_num,surplus_num,in_class_num,created_at,updated_at', '{\"class_id\":\"10010710\",\"continual_surplus_num\":\"14\",\"surplus_num\":\"14\",\"in_class_num\":\"0\",\"created_at\":\"2020-11-20 11:33:37\",\"updated_at\":\"2020-11-20 11:33:37\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167859, 'eas_test', 'classes_logs', '5fb78c5bab209-1605864539', 1, 'id,class_id,uid,content,type,created_at', 'id,class_id,uid,content,type,created_at', '{\"id\":\"69212\",\"class_id\":\"10010710\",\"uid\":\"121\",\"content\":\"新建待售班\",\"type\":\"1\",\"created_at\":\"2020-11-20 11:33:37\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167860, 'eas_test', 'class_position_exts', '5fb78c5bae38b-1605864539', 1, 'id,class_id,spare_position_num,total_position_num,position_detail,created_at,updated_at,deleted_at,deleted', 'id,class_id,spare_position_num,total_position_num,position_detail,created_at,updated_at,deleted_at,deleted', '{\"id\":\"59942\",\"class_id\":\"10010710\",\"spare_position_num\":\"14\",\"total_position_num\":\"14\",\"position_detail\":\"{\\\"1\\\":{\\\"id\\\":230474,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"2\\\":{\\\"id\\\":230475,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"3\\\":{\\\"id\\\":230476,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"4\\\":{\\\"id\\\":230477,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"5\\\":{\\\"id\\\":230478,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"6\\\":{\\\"id\\\":230479,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"7\\\":{\\\"id\\\":230480,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"8\\\":{\\\"id\\\":230481,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"9\\\":{\\\"id\\\":230482,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"10\\\":{\\\"id\\\":230483,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"11\\\":{\\\"id\\\":230484,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"12\\\":{\\\"id\\\":230485,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"13\\\":{\\\"id\\\":230486,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]},\\\"14\\\":{\\\"id\\\":230487,\\\"details\\\":[40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]}}\",\"created_at\":\"2020-11-20 11:33:37\",\"updated_at\":\"2020-11-20 11:33:37\",\"deleted_at\":\"\",\"deleted\":\"0\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167861, 'eas_test', 'schedule_event_log_202011', '5fb78c5bb2087-1605864539', 1, 'id,user_id,tag,curr_date,uuid,event,content,created_at,updated_at,sys_update_dc', 'id,user_id,tag,curr_date,uuid,event,content,created_at,updated_at,sys_update_dc', '{\"id\":\"3270\",\"user_id\":\"121\",\"tag\":\"10010710\",\"curr_date\":\"2020-11-19\",\"uuid\":\"2c78189a2ae111ebbf6400163e328457\",\"event\":\"schedule\",\"content\":\"{\\\"classes\\\":{\\\"id\\\":10010710,\\\"uid\\\":121,\\\"campus_id\\\":37,\\\"teaching_center_id\\\":13,\\\"product_id\\\":263,\\\"classroom_id\\\":0,\\\"classroom_ids\\\":\\\"21,21,0\\\",\\\"time_template_id\\\":56,\\\"name\\\":\\\"TEST-CULTURE-1120-01\\\",\\\"dict_class_status_id\\\":3,\\\"plan_min_class_number\\\":12,\\\"plan_max_class_number\\\":14,\\\"insert_class\\\":0,\\\"promoted_status\\\":0,\\\"coursed\\\":0,\\\"sold_degrees\\\":0,\\\"class_teacher_master_name\\\":null,\\\"guidance_teacher\\\":null,\\\"start_sale_time\\\":null,\\\"start_attend_class_time\\\":\\\"2020-11-20 00:00:00\\\",\\\"end_class_time\\\":null,\\\"inside_sell_time\\\":null,\\\"external_sell_time\\\":null,\\\"is_have_promoted\\\":0,\\\"is_original\\\":1,\\\"is_short_term\\\":0,\\\"is_online\\\":-1,\\\"is_demo\\\":0,\\\"is_make_up\\\":0,\\\"is_experiment\\\":0,\\\"is_demo_class\\\":0,\\\"demo_class_category_id\\\":0,\\\"set_online_status\\\":0,\\\"old_class_id\\\":null,\\\"half_start_date\\\":null,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"},\\\"class_teachers\\\":[{\\\"id\\\":89751,\\\"class_id\\\":10010710,\\\"teacher_id\\\":12026,\\\"dict_teacher_duty_id\\\":1,\\\"dict_teacher_attr_id\\\":1,\\\"teacher_sort\\\":1,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":null,\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"},{\\\"id\\\":89752,\\\"class_id\\\":10010710,\\\"teacher_id\\\":11877,\\\"dict_teacher_duty_id\\\":3,\\\"dict_teacher_attr_id\\\":2,\\\"teacher_sort\\\":1,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":null,\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"}],\\\"schedules\\\":[{\\\"id\\\":157882,\\\"class_id\\\":10010710,\\\"teacher_id\\\":12026,\\\"classroom_id\\\":21,\\\"time_model_type_id\\\":95,\\\"time_model_type_name\\\":\\\"\\\\u5e38\\\\u89c4\\\\u6a21\\\\u5f0f\\\",\\\"week_date\\\":3,\\\"flag_number\\\":1,\\\"uid\\\":121,\\\"start_date\\\":\\\"2020-11-20\\\",\\\"end_date\\\":\\\"2021-09-15\\\",\\\"start_time\\\":\\\"17:40:00\\\",\\\"end_time\\\":\\\"18:20:00\\\",\\\"is_foreign_class\\\":0,\\\"is_online\\\":0,\\\"is_dblt\\\":0,\\\"is_live\\\":0,\\\"is_centralized\\\":0,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":null,\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"},{\\\"id\\\":157883,\\\"class_id\\\":10010710,\\\"teacher_id\\\":12026,\\\"classroom_id\\\":21,\\\"time_model_type_id\\\":95,\\\"time_model_type_name\\\":\\\"\\\\u5e38\\\\u89c4\\\\u6a21\\\\u5f0f\\\",\\\"week_date\\\":3,\\\"flag_number\\\":2,\\\"uid\\\":121,\\\"start_date\\\":\\\"2020-11-20\\\",\\\"end_date\\\":\\\"2021-09-15\\\",\\\"start_time\\\":\\\"18:30:00\\\",\\\"end_time\\\":\\\"19:10:00\\\",\\\"is_foreign_class\\\":0,\\\"is_online\\\":0,\\\"is_dblt\\\":0,\\\"is_live\\\":0,\\\"is_centralized\\\":0,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":null,\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"},{\\\"id\\\":157884,\\\"class_id\\\":10010710,\\\"teacher_id\\\":12026,\\\"classroom_id\\\":0,\\\"time_model_type_id\\\":95,\\\"time_model_type_name\\\":\\\"\\\\u5e38\\\\u89c4\\\\u6a21\\\\u5f0f\\\",\\\"week_date\\\":6,\\\"flag_number\\\":0,\\\"uid\\\":121,\\\"start_date\\\":\\\"2020-11-20\\\",\\\"end_date\\\":\\\"2021-09-15\\\",\\\"start_time\\\":\\\"19:30:00\\\",\\\"end_time\\\":\\\"19:55:00\\\",\\\"is_foreign_class\\\":0,\\\"is_online\\\":1,\\\"is_dblt\\\":0,\\\"is_live\\\":0,\\\"is_centralized\\\":0,\\\"deleted\\\":0,\\\"deleted_at\\\":null,\\\"created_at\\\":\\\"2020-11-20 11:33:37\\\",\\\"updated_at\\\":null,\\\"sys_update_dc\\\":\\\"2020-11-20 11:33:37\\\"}]}\",\"created_at\":\"2020-11-20 11:33:40\",\"updated_at\":\"2020-11-20 11:33:40\",\"sys_update_dc\":\"2020-11-20 11:33:40\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167862, 'eas_test', 'classes', '5fb78c5bb601d-1605864539', 1, 'id,uid,campus_id,teaching_center_id,product_id,classroom_id,classroom_ids,time_template_id,name,dict_class_status_id,plan_min_class_number,plan_max_class_number,insert_class,promoted_status,coursed,sold_degrees,class_teacher_master_name,guidance_teacher,start_sale_time,start_attend_class_time,end_class_time,inside_sell_time,external_sell_time,is_have_promoted,is_original,is_short_term,is_online,is_demo,is_make_up,is_experiment,is_demo_class,demo_class_category_id,set_online_status,old_class_id,half_start_date,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,uid,campus_id,teaching_center_id,product_id,classroom_id,classroom_ids,time_template_id,name,dict_class_status_id,plan_min_class_number,plan_max_class_number,insert_class,promoted_status,coursed,sold_degrees,class_teacher_master_name,guidance_teacher,start_sale_time,start_attend_class_time,end_class_time,inside_sell_time,external_sell_time,is_have_promoted,is_original,is_short_term,is_online,is_demo,is_make_up,is_experiment,is_demo_class,demo_class_category_id,set_online_status,old_class_id,half_start_date,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"10010711\",\"uid\":\"121\",\"campus_id\":\"37\",\"teaching_center_id\":\"13\",\"product_id\":\"264\",\"classroom_id\":\"0\",\"classroom_ids\":\"5,5,0\",\"time_template_id\":\"81\",\"name\":\"TEST-CULTURE-1120-02\",\"dict_class_status_id\":\"3\",\"plan_min_class_number\":\"12\",\"plan_max_class_number\":\"14\",\"insert_class\":\"0\",\"promoted_status\":\"0\",\"coursed\":\"0\",\"sold_degrees\":\"0\",\"class_teacher_master_name\":\"\",\"guidance_teacher\":\"\",\"start_sale_time\":\"\",\"start_attend_class_time\":\"2020-11-24 00:00:00\",\"end_class_time\":\"\",\"inside_sell_time\":\"\",\"external_sell_time\":\"\",\"is_have_promoted\":\"0\",\"is_original\":\"1\",\"is_short_term\":\"0\",\"is_online\":\"-1\",\"is_demo\":\"0\",\"is_make_up\":\"0\",\"is_experiment\":\"0\",\"is_demo_class\":\"0\",\"demo_class_category_id\":\"0\",\"set_online_status\":\"0\",\"old_class_id\":\"\",\"half_start_date\":\"\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"2020-11-20 11:38:43\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167863, 'eas_test', 'class_teachers', '5fb78c5bb601d-1605864539', 1, 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"89753\",\"class_id\":\"10010711\",\"teacher_id\":\"12136\",\"dict_teacher_duty_id\":\"1\",\"dict_teacher_attr_id\":\"1\",\"teacher_sort\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167864, 'eas_test', 'class_teachers', '5fb78c5bb601d-1605864539', 1, 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,dict_teacher_duty_id,dict_teacher_attr_id,teacher_sort,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"89754\",\"class_id\":\"10010711\",\"teacher_id\":\"2889\",\"dict_teacher_duty_id\":\"3\",\"dict_teacher_attr_id\":\"2\",\"teacher_sort\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167865, 'eas_test', 'schedules', '5fb78c5bb601d-1605864539', 1, 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"157885\",\"class_id\":\"10010711\",\"teacher_id\":\"12136\",\"classroom_id\":\"5\",\"time_model_type_id\":\"141\",\"time_model_type_name\":\"常规模式\",\"week_date\":\"3\",\"flag_number\":\"1\",\"uid\":\"121\",\"start_date\":\"2020-11-24\",\"end_date\":\"2021-09-19\",\"start_time\":\"16:00:00\",\"end_time\":\"16:40:00\",\"is_foreign_class\":\"0\",\"is_online\":\"0\",\"is_dblt\":\"0\",\"is_live\":\"0\",\"is_centralized\":\"0\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167866, 'eas_test', 'schedules', '5fb78c5bb601d-1605864539', 1, 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"157886\",\"class_id\":\"10010711\",\"teacher_id\":\"12136\",\"classroom_id\":\"5\",\"time_model_type_id\":\"141\",\"time_model_type_name\":\"常规模式\",\"week_date\":\"3\",\"flag_number\":\"2\",\"uid\":\"121\",\"start_date\":\"2020-11-24\",\"end_date\":\"2021-09-19\",\"start_time\":\"15:00:00\",\"end_time\":\"15:40:00\",\"is_foreign_class\":\"0\",\"is_online\":\"0\",\"is_dblt\":\"0\",\"is_live\":\"0\",\"is_centralized\":\"0\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167867, 'eas_test', 'schedules', '5fb78c5bb601d-1605864539', 1, 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,teacher_id,classroom_id,time_model_type_id,time_model_type_name,week_date,flag_number,uid,start_date,end_date,start_time,end_time,is_foreign_class,is_online,is_dblt,is_live,is_centralized,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"157887\",\"class_id\":\"10010711\",\"teacher_id\":\"12136\",\"classroom_id\":\"0\",\"time_model_type_id\":\"141\",\"time_model_type_name\":\"常规模式\",\"week_date\":\"2\",\"flag_number\":\"0\",\"uid\":\"121\",\"start_date\":\"2020-11-24\",\"end_date\":\"2021-09-19\",\"start_time\":\"19:30:00\",\"end_time\":\"19:55:00\",\"is_foreign_class\":\"0\",\"is_online\":\"1\",\"is_dblt\":\"0\",\"is_live\":\"0\",\"is_centralized\":\"0\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (167868, 'eas_test', 'class_positions_37', '5fb78c5bb601d-1605864539', 1, 'id,class_id,degree_no,can_sell_num,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', 'id,class_id,degree_no,can_sell_num,status,deleted,deleted_at,created_at,updated_at,sys_update_dc', '{\"id\":\"230488\",\"class_id\":\"10010711\",\"degree_no\":\"1\",\"can_sell_num\":\"40\",\"status\":\"1\",\"deleted\":\"0\",\"deleted_at\":\"\",\"created_at\":\"2020-11-20 11:38:43\",\"updated_at\":\"2020-11-20 11:38:43\",\"sys_update_dc\":\"2020-11-20 11:38:43\"}', NULL, NULL, 0, NULL, '2020-11-20 17:28:59', '2020-11-20 17:28:59');
INSERT INTO `ddd_event_stream_copy1` VALUES (173685, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'name', '{\"name\":\"ss\"}', '', '', 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `ddd_event_stream_copy1` VALUES (174962, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"sss这是一个活动\",\"updated_at\":\"2020-12-24 17:34:45\"}', '', '', 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `ddd_event_stream_copy1` VALUES (175338, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'name,updated_at', '{\"name\":\"aabbccc\",\"updated_at\":\"2020-12-28 11:34:02\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-28 11:34:02', '2020-12-28 11:34:02');
INSERT INTO `ddd_event_stream_copy1` VALUES (175339, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'name,updated_at', '{\"name\":\"aabbddd\",\"updated_at\":\"2020-12-28 11:35:24\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-28 11:35:24', '2020-12-28 11:35:24');
INSERT INTO `ddd_event_stream_copy1` VALUES (175340, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'name,updated_at', '{\"name\":\"aabbttt\",\"updated_at\":\"2020-12-28 11:37:22\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-28 11:37:22', '2020-12-28 11:37:22');
INSERT INTO `ddd_event_stream_copy1` VALUES (175341, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'name,updated_at', '{\"name\":\"aabbttsssst\",\"updated_at\":\"2020-12-28 13:22:39\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-28 13:22:39', '2020-12-28 13:22:39');
INSERT INTO `ddd_event_stream_copy1` VALUES (175342, 'codeper', 'user', '', 2, 'id,name,address,age,test,created_at,updated_at', 'address,updated_at', '{\"address\":\"苦苦sa\",\"updated_at\":\"2020-12-28 14:14:07\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-28 14:14:08', '2020-12-28 14:14:08');
INSERT INTO `ddd_event_stream_copy1` VALUES (175343, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"aaatrre\",\"updated_at\":\"2020-12-29 12:09:13\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 12:09:14', '2020-12-29 12:09:14');
INSERT INTO `ddd_event_stream_copy1` VALUES (175344, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"oooooooxxx\",\"updated_at\":\"2020-12-29 12:20:01\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 12:20:03', '2020-12-29 12:20:03');
INSERT INTO `ddd_event_stream_copy1` VALUES (175345, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"ssoooooooxxx\",\"updated_at\":\"2020-12-29 12:21:46\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 12:21:48', '2020-12-29 12:21:48');
INSERT INTO `ddd_event_stream_copy1` VALUES (175346, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"rrrrrrrttrrrrr\",\"updated_at\":\"2020-12-29 15:00:59\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:01:00', '2020-12-29 15:01:00');
INSERT INTO `ddd_event_stream_copy1` VALUES (175347, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"s\",\"updated_at\":\"2020-12-29 15:01:16\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:01:18', '2020-12-29 15:01:18');
INSERT INTO `ddd_event_stream_copy1` VALUES (175348, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"wwwwww\",\"updated_at\":\"2020-12-29 15:01:22\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:01:24', '2020-12-29 15:01:24');
INSERT INTO `ddd_event_stream_copy1` VALUES (175349, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"uuuuuu\",\"updated_at\":\"2020-12-29 15:01:24\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:01:26', '2020-12-29 15:01:26');
INSERT INTO `ddd_event_stream_copy1` VALUES (175350, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'content,updated_at', '{\"content\":\"c2Zkc2Y=\",\"updated_at\":\"2020-12-29 15:01:44\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:01:46', '2020-12-29 15:01:46');
INSERT INTO `ddd_event_stream_copy1` VALUES (175351, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'content,updated_at', '{\"content\":\"ssss\",\"updated_at\":\"2020-12-29 15:28:34\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:28:35', '2020-12-29 15:28:35');
INSERT INTO `ddd_event_stream_copy1` VALUES (175352, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'member_id,updated_at', '{\"member_id\":\"%!s(int32=100)\",\"updated_at\":\"2020-12-29 15:29:10\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:29:12', '2020-12-29 15:29:12');
INSERT INTO `ddd_event_stream_copy1` VALUES (175353, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'member_id,updated_at', '{\"member_id\":\"30000\",\"updated_at\":\"2020-12-29 15:43:49\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:43:51', '2020-12-29 15:43:51');
INSERT INTO `ddd_event_stream_copy1` VALUES (175354, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'content,updated_at', '{\"content\":\"www\",\"updated_at\":\"2020-12-29 15:43:54\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:43:56', '2020-12-29 15:43:56');
INSERT INTO `ddd_event_stream_copy1` VALUES (175355, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'content,updated_at', '{\"content\":\"rrrrrrr\",\"updated_at\":\"2020-12-29 15:43:56\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 15:43:58', '2020-12-29 15:43:58');
INSERT INTO `ddd_event_stream_copy1` VALUES (175356, 'codeper', 'activity', '', 2, 'activity_id,activity_pic,title,member_id,start_at,attend_num,content,status,is_deleted,deleted_at,created_at,updated_at', 'title,updated_at', '{\"title\":\"qqqqqqq\",\"updated_at\":\"2020-12-29 16:07:33\"}', '', '', 0, '0000-00-00 00:00:00', '2020-12-29 16:07:35', '2020-12-29 16:07:35');
COMMIT;

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
  `status` tinyint(4) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=234 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (200, '菜单列表', 0, 0, 0, '/admin/menu/list.html', 'fa-bars', 0, '2020-02-16 09:14:38', '2020-12-31 11:06:42', 0, '2020-07-21 13:09:05');
INSERT INTO `menu` VALUES (201, '管理员列表', 0, 0, 0, '/admin/admins/list.html', 'fa-user', 0, '2020-02-16 09:14:38', '2020-12-31 11:06:51', 0, NULL);
INSERT INTO `menu` VALUES (203, '资源列表', 0, 0, 0, '/admin/resource/list.html', 'fa-tag', 0, '2020-02-16 09:14:38', '2020-12-31 11:06:54', 0, '2020-11-15 22:48:13');
INSERT INTO `menu` VALUES (205, '权限管理', 0, 0, 0, '', 'fa-share', 0, '2020-02-16 09:14:38', '2020-12-31 11:07:01', 0, NULL);
INSERT INTO `menu` VALUES (206, '权限节点列表', 205, 0, 0, '/admin/permissions/list.html', 'fa-tag', 1, '2020-02-16 09:14:38', '2020-07-22 13:48:33', 0, NULL);
INSERT INTO `menu` VALUES (207, '角色列表', 205, 0, 0, '/admin/roles/list.html', 'fa-tag', 1, '2020-02-16 09:14:38', '2020-07-22 13:48:36', 0, '2020-07-20 22:42:33');
INSERT INTO `menu` VALUES (231, '业务事件数据模型', 0, 0, 0, '', '', 1, '2020-11-12 18:11:19', '2020-11-15 20:38:55', 0, NULL);
INSERT INTO `menu` VALUES (232, '事件列表', 231, 0, 0, '/admin/event/list.html', '', 1, '2020-11-12 18:11:52', '2020-11-12 18:14:26', 0, NULL);
INSERT INTO `menu` VALUES (233, 'binlog数据流', 231, 0, 0, '/admin/event/stream_list.html', '', 1, '2020-11-12 18:26:20', '2020-11-15 20:41:49', 0, NULL);
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

SET FOREIGN_KEY_CHECKS = 1;
