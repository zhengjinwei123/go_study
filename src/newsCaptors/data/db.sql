
CREATE TABLE IF NOT EXISTS  `t_option` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `t_option` */

insert  into `t_option`(`id`,`name`,`value`) values (1,'sitename','捕获者'),(2,'siteurl','http://zhengjinwei.top:3006'),(3,'subtitle','捕获者'),(4,'pagesize','10'),(5,'keywords','Go语言,beego,爬虫'),(6,'description','基于GO语言的网络爬虫项目'),(7,'email','2538698032@qq.com'),(8,'theme','default'),(9,'timezone','8'),(10,'stat','');


CREATE TABLE IF NOT EXISTS `t_hobby` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `hobby` varchar(255) NOT NULL DEFAULT '' COMMENT '爱好',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '分类',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_hobby_type` (`hobby`,`type`),
  KEY `idx_type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

CREATE TABLE  IF NOT EXISTS `t_captorinfo` (
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
) ENGINE=InnoDB AUTO_INCREMENT=6647 DEFAULT CHARSET=utf8;