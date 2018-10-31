package main

import (
	_ "mygoblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"beegoProj/models"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	//"beegoProj/utils"
	//"time"
	"time"
)

func initRedis(){
	redis,err := cache.NewCache("redis",`{"conn":"106.15.195.250:6379","password":"123456","dbNum":"0"}`)
	if err != nil{
		panic(err)
	}
	redis.Put("test","hello",time.Second*100)
	v := redis.Get("test")
	fmt.Println("redis key:test:value:" + cache.GetString(v))
}

func initMysql() {
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
		//Id: 100,
		//UserName: "ZJW",
		//Password: utils.MD5("123456"),
		//Email: "2538698032@qq.com",
		//LastLogin: time.Now(),
		//LoginCount: 1,
		//LastIp: "127.0.0.1",
		//Authkey: "auth_key_1",
		//Active: 1,
	}
	//if err := user.Insert(); err != nil {
	//	panic(err)
	//}

	//user.UserName = "zjw1111"
	//if err := user.Update();err != nil {
	//	panic(err)
	//}

	user.Id = 100
	user.Read()
	fmt.Println(user)
}

func init(){
	initRedis()
	//initMysql()
}

func main() {
	beego.Run()
}

