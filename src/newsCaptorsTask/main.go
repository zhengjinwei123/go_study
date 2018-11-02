package main

import (
	"fmt"
	"newsCaptorsTask/global"
)


func main() {
	conf := global.AppConf()
	fmt.Println(conf.GetValue("mysql","port"))
	fmt.Println(conf.GetValue("base","appName"))
	fmt.Println(conf.GetValue("mysql","host"))
}
