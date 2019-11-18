DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `articles`;
DROP TABLE IF EXISTS `categories`;
DROP TABLE IF EXISTS `tags`;

CREATE TABLE `users`
(
  `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username`   varchar(255) NOT NULL UNIQUE,
  `password`   varchar(255) NOT NULL,
  `avatar`     varchar(255) NOT NULL,
  `created_at` timestamp NULL,
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `users` VALUES (0,'admin','admin','https://user-gold-cdn.xitu.io/2019/5/29/16b028263cf8b532?imageView2/1/w/100/h/100/q/85/format/webp/interlace/1','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);

CREATE TABLE `articles`
(
  `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title`       varchar(255) NOT NULL,
  `content`     varchar(255) NOT NULL,
  `category_id` int(11) UNSIGNED NOT NULL,
  `tag_id`      int(11) UNSIGNED NOT NULL,
  `user_id`     int(11) UNSIGNED NOT NULL,
  `created_at`  timestamp NULL,
  `updated_at`  timestamp NULL,
  `deleted_at`  timestamp NULL,
  PRIMARY KEY (`id`)
);
CREATE TABLE `categories`
(
  `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL UNIQUE,
  `created_at`    timestamp NULL,
  `updated_at`    timestamp NULL,
  `deleted_at`    timestamp NULL,
  PRIMARY KEY (`id`)
);
CREATE TABLE `tags`
(
  `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_name`   varchar(255) NOT NULL UNIQUE,
  `created_at` timestamp NULL,
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
     `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
     `name` varchar(50) DEFAULT '' COMMENT '名字',
     `path` varchar(50) DEFAULT '' COMMENT '访问路径',
     `method` varchar(50) DEFAULT '' COMMENT '资源请求方式',
     `created_at` timestamp NULL,
     `updated_at` timestamp NULL,
     `deleted_at` timestamp NULL,
     PRIMARY KEY (`id`)
)

-- INSERT INTO `menu` VALUES (1, '获取用户信息', '/v1/auth/info', 'GET', null, null);
-- INSERT INTO `menu` VALUES (2, '获取用户列表', '/v1/user', 'POST', null, null);
-- INSERT INTO `menu` VALUES (3, '获取单个用户', '/v1/user/:id', 'GET', null, null);


-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(50) DEFAULT '' COMMENT '名字',
     `created_at` timestamp NULL,
     `updated_at` timestamp NULL,
     `deleted_at` timestamp NULL,
     PRIMARY KEY (`id`)
)

-- ----------------------------
-- Records of role
-- ----------------------------
-- INSERT INTO `role` VALUES (1, 'admin', null, null);
-- INSERT INTO `role` VALUES (2, 'ant', null, null);


-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
    `menu_id` int(11) unsigned DEFAULT NULL COMMENT '菜单ID',
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`)
)

-- ----------------------------
-- Records of role_menu
-- ----------------------------
-- INSERT INTO `role_menu` VALUES (1, 2, 1);
-- INSERT INTO `role_menu` VALUES (2, 2, 2);
-- INSERT INTO `role_menu` VALUES (3, 2, 3);


-- ----------------------------
-- Table structure for users_role
-- ----------------------------
DROP TABLE IF EXISTS `users_role`;
CREATE TABLE `users_role` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(11) unsigned DEFAULT NULL COMMENT '用户ID',
    `role_id` int(11) unsigned DEFAULT NULL COMMENT '角色ID',
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`)
);

-- ----------------------------
-- Records of users_role
-- ----------------------------
-- INSERT INTO `users_role` VALUES (1, 1, 1);