/*

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 25/04/2020 11:55:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_admin_member
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin_member`;
CREATE TABLE `tb_admin_member` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `passwd` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
  `name` varchar(36) NOT NULL DEFAULT '' COMMENT '姓名',
  `role` char(8) NOT NULL DEFAULT '' COMMENT '角色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='管理员信息表';

-- ----------------------------
-- Records of tb_admin_member
-- ----------------------------
BEGIN;
INSERT INTO `tb_admin_member` VALUES (1, '2020-04-22 20:45:52', '2020-04-22 20:45:52', NULL, 'admin', '$2a$04$2POG0DEvZFa0mBc2izxb7uuvjSFjo.gr0GQca1qZPurlZxqmoCIK.', '', '');
COMMIT;

-- ----------------------------
-- Table structure for tb_admin_session
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin_session`;
CREATE TABLE `tb_admin_session` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `token` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `member_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `available_time` timestamp NULL DEFAULT NULL COMMENT '有效时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='管理员信息表';


-- ----------------------------
-- Table structure for tb_test_class
-- ----------------------------
DROP TABLE IF EXISTS `tb_test_class`;
CREATE TABLE `tb_test_class` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(36) NOT NULL DEFAULT '' COMMENT '班级名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='test-班级信息表';

-- ----------------------------
-- Records of tb_test_class
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_test_score
-- ----------------------------
DROP TABLE IF EXISTS `tb_test_score`;
CREATE TABLE `tb_test_score` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `student_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '学生ID，对应学生表ID',
  `course` varchar(20) NOT NULL DEFAULT '' COMMENT '考试科目',
  `score` float(3,2) NOT NULL DEFAULT '0.00' COMMENT '成绩',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='test-成绩信息表';

-- ----------------------------
-- Records of tb_test_score
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_test_student
-- ----------------------------
DROP TABLE IF EXISTS `tb_test_student`;
CREATE TABLE `tb_test_student` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(36) NOT NULL DEFAULT '' COMMENT '姓名',
  `number` char(20) NOT NULL DEFAULT '' COMMENT '学号',
  `gender` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '性别，1-男，2-女',
  `birthday` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  `class_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '班级ID，对应班级表ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='test-学生信息表';

-- ----------------------------
-- Records of tb_test_student
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
