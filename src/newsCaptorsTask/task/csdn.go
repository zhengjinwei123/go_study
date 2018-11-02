package task

import (
	"fmt"
	"newsCaptorsTask/config"
	"strings"
)

type CSDN struct {
	Base
}

func (this *CSDN) Init(){
	fmt.Println("csdn init")
	appConf := config.AppConf()
	csdnTaskList := strings.Split(appConf.GetValue("task",this.TaskName),"|")
	for _,v := range csdnTaskList {
		this.CronList = append(this.CronList, strings.TrimSpace(v))
	}
}

func (this *CSDN) Worker(){
	fmt.Println("csdn worker")
}