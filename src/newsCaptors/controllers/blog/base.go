package blog

import (
	"github.com/astaxie/beego"
	"goblog/util"
)

type baseController struct {
	beego.Controller
	moduleName string
	controllerName string
	actionName string
	options map[string]string
	cache *util.LruCache
}

func (this *baseController) Prepare(){

}