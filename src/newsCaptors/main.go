package main

import (
	"github.com/astaxie/beego"
	_ "newsCaptors/routers"
	"newsCaptors/util"
	"newsCaptors/models"
)

func init(){
	util.Factory.Set("cache", func() (interface{}, error) {
		mc := util.NewLRUCache(1000)
		return mc,nil
	})
	models.Init()
}

func main() {
	beego.Run()
}

