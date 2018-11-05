package task

import (
	"github.com/robfig/cron"
)

type ResultType struct {
	Desc    string
	Link    string
}

type Base struct {
	TaskName string
	CronList []string
	ResultList []*ResultType
}



func (this *Base) Init(){

}

func (this *Base) Run(worker func(keyWord string),keyWord string){
	for _,spec := range this.CronList {
		go func(spec string,that *Base){
			newCron := cron.New()
			newCron.AddFunc(spec, func() {
				worker(keyWord)
			})
			newCron.Start()
		}(spec,this)
	}
}

