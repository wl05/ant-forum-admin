DROP TABLE IF EXISTS `user`;
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
