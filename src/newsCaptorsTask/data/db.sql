use `newscaptor`
CREATE TABLE `t_captorinfo` (
  `id` bigint unsigned not null auto_increment,
  `url` varchar(255) not null default '' comment '链接地址',
  `createAt` datetime not null default CURRENT_TIMESTAMP,
  `updateAt` datetime not null default CURRENT_TIMESTAMP on update  CURRENT_TIMESTAMP,
  primary key (`id`),
  key `idx_createAt` (`createAt`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;