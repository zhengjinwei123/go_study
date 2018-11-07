use `newscaptor`
CREATE TABLE `t_captorinfo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
  `keyword` varchar(255) NOT NULL DEFAULT '' COMMENT '关键字',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_url_desc` (`url`,`desc`),
  KEY `idx_createAt` (`createAt`),
  KEY `idx_keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE `t_hobby` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `hobby` varchar(255) NOT NULL DEFAULT '' COMMENT '爱好',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '分类',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_hobby_type` (`hobby`,`type`),
  KEY `idx_type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;