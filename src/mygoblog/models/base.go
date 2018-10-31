package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init(){
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbname := beego.AppConfig.String("dbname")
	dbpassword := beego.AppConfig.String("dbpassword")

	if dbport == ""{
		dbport = "3306"
	}

	dsn := dbuser + ":" + "@tcp(" + dbhost+":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia/Shanghai"
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModel(new(User),new(Post),new(Tag),new(TagPost),new(Option))
}

func TableName(str string) string {
	return beego.AppConfig.String("deprefix") + str
}