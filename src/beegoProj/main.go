package main

import (
	_ "mygoblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"beegoProj/models"
	"beegoProj/utils"
	"time"
)


func initMysql(){
	orm.RegisterDriver("mysql",orm.DRMySQL)

	dbHost := beego.AppConfig.String("db_mysql_host")
	dbUserName := beego.AppConfig.String("db_user_name")
	dbPort := beego.AppConfig.String("db_mysql_port")
	dbPassword := beego.AppConfig.String("db_mysql_password")
	database := beego.AppConfig.String("db_mysql_database")

	url := dbUserName+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+database+"?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(url)

	err := orm.RegisterDataBase("default","mysql",url)
	if err != nil {
		panic(fmt.Sprintf("database register err:%x",err))
	}
	orm.RegisterModel(new(models.User))
	orm.RunCommand()

	user := &models.User{
		Id: 100,
		UserName: "ZJW",
		Password: utils.MD5("123456"),
		Email: "2538698032@qq.com",
		LastLogin: time.Now(),
		LoginCount: 1,
		LastIp: "127.0.0.1",
		Authkey: "auth_key_1",
		Active: 1,
	}
	if err := user.Insert(); err != nil {
		panic(err)
	}
}

func init(){
	initMysql()
}

func main() {
	beego.Run()
}

