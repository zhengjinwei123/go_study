package manager

import (
	"newsCaptorsTask/task"
	"net/url"
	"newsCaptorsTask/model"
	"newsCaptorsTask/config"
)

func init(){
	log := MyLogger("captor")

	log.Debug("task Manager init")
	csdnTask := new(task.CSDN)
	csdnTask.TaskName = "csdn"
	csdnTask.Init()

	hobby := new(model.Hobby)
	hobby.Type = config.EHOBBY_TYPE_CSDN

	var hobbyList []*model.Hobby
	t,err := hobby.Query().All(&hobbyList)
	if err == nil && t > 0 {
		for _,v := range hobbyList {
			hb := (*v).Hobby
			log.Debug("task process csdn task,hobby:",hb)
			csdnTask.Run(csdnTask.Worker,url.QueryEscape(hb))
		}
	}
}