package blog

import (
	"github.com/astaxie/beego"
	"newsCaptors/util"
	"strings"
)

type baseController struct {
	beego.Controller
	moduleName string
	controllerName string
	actionName string
	//options map[string]string
	cache *util.LRUCache
}

func (this *baseController) Prepare(){
	controllerName,actionName := this.GetControllerAndAction()
	this.moduleName = "blog"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	cache, _ := util.Factory.Get("cache")
	this.cache = cache.(*util.LRUCache)
}