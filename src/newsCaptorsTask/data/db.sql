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
) ENGINE=InnoDB AUTO_INCREMENT=1230 DEFAULT CHARSET=utf8