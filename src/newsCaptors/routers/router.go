package routers

import (
	"newsCaptors/controllers/blog"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&blog.IndexController{},"*:Index")
}
