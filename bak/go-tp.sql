SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_power
-- ----------------------------
DROP TABLE IF EXISTS `admin_power`;
CREATE TABLE `admin_power` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL COMMENT '权限名字',
  `controller` varchar(100) DEFAULT NULL COMMENT '控制器',
  `action` varchar(100) DEFAULT NULL COMMENT '方法名',
  `active` varchar(100) DEFAULT NULL COMMENT '父级方法',
  `level` int(5) DEFAULT '0' COMMENT '等级',
  `sort` int(3) DEFAULT '0' COMMENT '排序',
  `seo_description` varchar(500) DEFAULT NULL COMMENT 'seo描述',
  `seo_title` varchar(100) DEFAULT NULL COMMENT 'seo标题',
  `seo_keyword` varchar(200) DEFAULT NULL COMMENT 'seo关键词',
  `is_show` tinyint(1) DEFAULT '0' COMMENT '菜单1显示 0不显示',
  `ico` varchar(30) DEFAULT NULL COMMENT 'fa ico图标',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常 0删除',
  `group` varchar(50) DEFAULT NULL COMMENT '组',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin_power
-- ----------------------------
BEGIN;
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (1, '系统设置', 'System', '', '', 0, 0, NULL, NULL, NULL, 1, 'fa-cogs', NULL, 1571211001, 1, 'wx');
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (2, '权限设置', 'System', 'power', 'power', 1, 0, NULL, NULL, NULL, 1, '', 1567144226, 1571665933, 1, 'wx');
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (3, '角色管理', 'System', 'role', 'role', 1, 0, NULL, NULL, NULL, 1, NULL, NULL, 1571665929, 1, 'wx');
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (4, '用户管理', 'System', 'user', 'user', 1, 0, NULL, NULL, NULL, 1, NULL, NULL, 1571665928, 1, 'wx');
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (5, '控制面板', 'Index', '', '', 0, 99, NULL, NULL, NULL, 1, 'fa-dashboard', NULL, 1571665928, 1, 'wx');
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (23, '资讯', 'Basis', '', '', 0, 0, NULL, NULL, NULL, 1, 'fa-database', 1614592695, 1614592695, 1, NULL);
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (24, '文章列表', 'Basis', 'article', 'index', 23, 0, NULL, NULL, NULL, 1, '', 1614592808, 1614592808, 1, NULL);
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (25, '文章编辑', 'Basis', 'editarticle', 'index', 23, 0, NULL, NULL, NULL, 0, '', 1614654075, 1614654075, 1, NULL);
INSERT INTO `admin_power` (`id`, `title`, `controller`, `action`, `active`, `level`, `sort`, `seo_description`, `seo_title`, `seo_keyword`, `is_show`, `ico`, `created_at`, `updated_at`, `status`, `group`) VALUES (26, '主题列表', 'Basis', 'topic', 'index', 23, 0, NULL, NULL, NULL, 1, '', 1614592889, 1614592889, 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200) DEFAULT NULL COMMENT '权限标题',
  `description` varchar(300) DEFAULT NULL COMMENT '权限描述',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `status` int(1) DEFAULT '1' COMMENT '状态1正常 0删除',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `config` text COMMENT '配置',
  `group` varchar(50) DEFAULT NULL COMMENT '群组',
  `superior` int(3) DEFAULT '0' COMMENT '上级',
  `visit_team` varchar(200) NOT NULL COMMENT '所能访问的其他组',
  `operate` tinyint(1) DEFAULT '1' COMMENT '公众号操作权限',
  `notice` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_role` (`id`, `title`, `description`, `created_at`, `updated_at`, `status`, `sort`, `config`, `group`, `superior`, `visit_team`, `operate`, `notice`) VALUES (1, '超级管理员', '全部权限', NULL, 1567394906, 1, 99, '0', NULL, -1, '', 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(50) DEFAULT NULL COMMENT '密码',
  `name` varchar(11) DEFAULT NULL COMMENT '备注',
  `salt` int(6) DEFAULT NULL COMMENT '安全验证码',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
  `role_id` int(3) DEFAULT NULL COMMENT '所属权限',
  `status` int(2) DEFAULT '1' COMMENT '状态1正常 0删除',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `config` text COMMENT '服务号',
  `code` varchar(8) DEFAULT 't',
  `code1` varchar(8) DEFAULT 'r',
  `config1` text COMMENT '开放者平台2',
  `config2` text COMMENT '个人号',
  `code2` varchar(8) DEFAULT NULL,
  `wxappid` varchar(255) DEFAULT '' COMMENT '小程序appid',
  `team` int(3) DEFAULT '0' COMMENT '所属团队',
  `domain` varchar(200) DEFAULT NULL COMMENT '地址域名',
  `domain2` varchar(200) DEFAULT NULL COMMENT '替换地址域名',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`username`,`password`,`name`,`salt`,`role_id`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
BEGIN;
INSERT INTO `admin_user` (`id`, `username`, `password`, `name`, `salt`, `email`, `phone`, `role_id`, `status`, `created_at`, `updated_at`, `config`, `code`, `code1`, `config1`, `config2`, `code2`, `wxappid`, `team`, `domain`, `domain2`) VALUES (1, 'admin', '0079ac91935392eee406647cc43b98eb', '总管理', 454974, NULL, NULL, 1, 1, 1606963190, 1606963190, '0', 't', 'r', '0', NULL, 'c', '', 0, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` int(1) DEFAULT NULL COMMENT '状态',
  `topic_id` int(11) DEFAULT NULL COMMENT '栏目id',
  `title` varchar(300) COLLATE utf8mb4_bin DEFAULT NULL,
  `content` text COLLATE utf8mb4_bin,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `litpic` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (26, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (27, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (28, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (29, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (30, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (31, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (32, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
INSERT INTO `article` (`id`, `status`, `topic_id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, `litpic`) VALUES (33, 1, NULL, '文章1', '正文1', 1664378735, 1664378735, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(300) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `sort` int(11) DEFAULT NULL,
  `status` int(1) DEFAULT NULL COMMENT '0删除 1正常',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of topic
-- ----------------------------
BEGIN;
INSERT INTO `topic` (`id`, `title`, `sort`, `status`, `created_at`, `updated_at`) VALUES (1, '主机游戏', 1, 1, 1614654838, 1614654838);
INSERT INTO `topic` (`id`, `title`, `sort`, `status`, `created_at`, `updated_at`) VALUES (2, '电脑游戏', 2, 1, 1614654846, 1665130556);
INSERT INTO `topic` (`id`, `title`, `sort`, `status`, `created_at`, `updated_at`) VALUES (3, '手机游戏', 3, 1, 1614654852, 1614654852);

-- ----------------------------
-- Table structure for upload
-- ----------------------------
DROP TABLE IF EXISTS `upload`;
CREATE TABLE `upload` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `uid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of upload
-- ----------------------------
BEGIN;
INSERT INTO `upload` (`id`, `url`, `created_at`, `uid`) VALUES (1, '/uploads/20210302/09d0d286578b2a2c984644c05f74a47d.jpeg', 1614654663, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
