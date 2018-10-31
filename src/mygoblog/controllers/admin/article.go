package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"mygoblog/models"
	"mygoblog/util"
	"os"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) List() {
	var (
		page       int
		pagesize   int = 10
		status     int
		offset     int
		list       []*models.Post
		post       models.Post
		searchtype string
		keyword    string
	)

	searchtype = this.GetString("searchtype")
	keyword = this.GetString("keyword")
	status, _ = this.GetInt("status")
	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	query := post.Query().Filter("status", status)

	if keyword != "" {
		switch searchtype {
		case "title":
			query = query.Filter("title__icontains", keyword)
		case "author":
			query = query.Filter("author__icontains", keyword)
		case "tag":
			query = query.Filter("tags__icontains", keyword)
		}
	}

	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-is_top", "-post_time").Limit(pagesize, offset).All(&list)
	}

	this.Data["searchtype"] = searchtype
	this.Data["keyword"] = keyword
	this.Data["count_1"], _ = post.Query().Filter("status", 1).Count()
	this.Data["count_2"], _ = post.Query().Filter("status", 2).Count()
	this.Data["status"] = status
	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize, fmt.Sprintf("/admin/article/list?status=%d&searchtype=%s&keyword=%s", status, searchtype, keyword), true).ToString()
	this.display()
}
