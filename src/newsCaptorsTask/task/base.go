package task

import (
	"github.com/robfig/cron"
	"fmt"
)

type Base struct {
	TaskName string
	CronList []string
}


func (this *Base) Init(){

}

func (this *Base) Run(){
	for _,spec := range this.CronList {
		fmt.Println(spec)
		go func(spec string,that *Base){
			newCron := cron.New()
			newCron.AddFunc(spec, func() {
				that.Worker()
			})
			newCron.Start()
		}(spec,this)
	}
}

func (this *Base) Worker(){
	fmt.Println("base worker")
}

