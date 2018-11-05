package task

import (
	"newsCaptorsTask/config"
	"strings"
	"log"
	"newsCaptorsTask/util/spider"
	"regexp"
	"strconv"
	"newsCaptorsTask/model"
	"fmt"
	"time"
	"net/url"
)

type CSDN struct {
	Base
	PageNum int
}

func (this *CSDN) Init(){
	log.Println("csdn init")
	appConf := config.AppConf()
	csdnTaskList := strings.Split(appConf.GetValue("task",this.TaskName),"|")
	for _,v := range csdnTaskList {
		this.CronList = append(this.CronList, strings.TrimSpace(v))
	}
	this.PageNum = 1
}


func (this *CSDN) Worker(keyWord string){
	spy := spider.Spider {
		Url: "https://so.csdn.net/so/search/s.do?q="+keyWord+"&t=%20&u=%s",
		PageContent: "",
	}

	spy.MatchFunc = func() int {
		exp1 := regexp.MustCompile(`<dl class="search-list J_search">\r\n+\s+<dt>\r\n+\s+<a href="(.*?)" target="_blank" strategy="SearchFromCsdn">(.*?)</a>`)
		result1 := exp1.FindAllStringSubmatch(spy.PageContent,-1)

		log.Println("match func called:"+strconv.Itoa(this.PageNum))
		for _,v :=range result1 {
			res := new(ResultType)
			for key,value := range v {
				if key == 0 {
					continue
				}
				if key == 1 {
					res.Link = value
				} else {
					res.Desc = strings.Replace(value,"<em>","",-1)
					res.Desc = strings.Replace(res.Desc,"</em>","",-1)
				}
			}
			this.ResultList = append(this.ResultList,res)
		}
		for _,v := range this.ResultList {
			// 数据入库
			log.Println("haha",v.Desc,v.Link)
		}

		if len(result1) > 0 && this.PageNum <= 20 {
			this.PageNum ++
			url := spy.Url + "&p=" + strconv.Itoa(this.PageNum)
			spy.GetPage(url)
			spy.ProcessPage()
		} else {
			// 入库
			this.InsertToDB(keyWord)
		}
		return 0
	}

	spy.GetPage(spy.Url)
	spy.ProcessPage()
}

func (this *CSDN) InsertToDB(keyWord string){
	keyWord,_ = url.QueryUnescape(keyWord)
	for _,v := range this.ResultList{
		m := &model.CaptorInfo{
			Url: v.Link,
			Desc: v.Desc,
			Keyword: keyWord,
			Createat:time.Now(),
			Updateat:time.Now(),
		}
		err := m.Insert()
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Println("insert to db success")
		}
	}
	fmt.Println("insert to db end")
}