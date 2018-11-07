package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func Init() {
	host := beego.AppConfig.String("dbhost")
	port := beego.AppConfig.String("dbport")
	user := beego.AppConfig.String("dbuser")
	password := beego.AppConfig.String("dbpassword")
	database := beego.AppConfig.String("dbdatabase")

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