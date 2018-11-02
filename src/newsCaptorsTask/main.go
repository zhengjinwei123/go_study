package main

import (
	"newsCaptorsTask/util/spider"
	"newsCaptorsTask/config"
	_ "newsCaptorsTask/global"
	"time"
)


func main() {
	config.AppConf()
	spy := spider.Spider {
		Url: "http://www.baidu.com",
		Reg: "",
		PageContent:"",
		ResultList: make([]string,1000),
	}
	spy.GetPage()
	//fmt.Println(spy.PageContent)
	time.Sleep(time.Hour)
}
