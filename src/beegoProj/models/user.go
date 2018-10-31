package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/**
CREATE TABLE `t_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(15) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `login_count` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_ip` varchar(15) NOT NULL DEFAULT '0' COMMENT '最后登录ip',
  `last_login` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `authkey` char(10) NOT NULL DEFAULT '' COMMENT '登录key',
  `active` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否激活',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8
 */
type User struct {
	Id int
	UserName string `orm:"unique;size(15)"`
	Password string `orm:"size(32)"`
	Email string `orm:"size(50)"`
	LastLogin time.Time `orm:"auto_now_add;type(datetime)"`
	LoginCount int
	LastIp string `orm:"size(32)"`
	Authkey string `orm:"size(10)"`
	Active int8
}

func (this *User) TableName() string {
	return "t_user"
}

func (this *User) Insert() error {
	if _,err := orm.NewOrm().Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(this,fields...);err != nil {
		return err
	}
	return nil
}

func (this *User) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(this,fields...);err != nil {
		return err
	}
	return nil
}

func (this *User) Delete() error {
	if _,err := orm.NewOrm().Delete(this);err != nil {
		return err
	}
	return nil
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}