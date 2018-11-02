package main

import (
	"fmt"
	"newsCaptorsTask/global"
	//"github.com/robfig/cron"
	//"time"
)


func main() {
	conf := global.AppConf()
	fmt.Println(conf.GetValue("mysql","port"))
	fmt.Println(conf.GetValue("base","appName"))
	fmt.Println(conf.GetValue("mysql","host"))

	//c := cron.New()
	//spec := "*/5 * * * * ?"
	//i := 0
	//c.AddFunc(spec,func(){
	//	i++
	//	fmt.Println("cron running:",i)
	//})
	//c.Start()
	//for i != -1 {
	//
	//}
}
