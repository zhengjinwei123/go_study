package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)


func Init() {
	host := beego.AppConfig.String("mysql::host")
	port := beego.AppConfig.String("mysql::port")
	user := beego.AppConfig.String("mysql::user")
	password := beego.AppConfig.String("mysql::password")
	database := beego.AppConfig.String("mysql::database")

	orm.RegisterDriver("mysql",orm.DRMySQL)
	url := user + ":" + password + "@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local"

	err := orm.RegisterDataBase("default","mysql",url)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(CaptorInfo))
	orm.RunCommand()
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}