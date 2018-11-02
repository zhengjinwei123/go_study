package task

import (
	"github.com/robfig/cron"
)

type Base struct {
	TaskName string
	CronList []string
}


func (this *Base) Init(){

}

func (this *Base) Run(worker func()){
	for _,spec := range this.CronList {
		go func(spec string,that *Base){
			newCron := cron.New()
			newCron.AddFunc(spec, func() {
				worker()
			})
			newCron.Start()
		}(spec,this)
	}
}

