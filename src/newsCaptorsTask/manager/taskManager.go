package manager

import (
	"newsCaptorsTask/task"
	"log"
	"net/url"
)

func init(){
	log.Println("task Manager init")
	csdnTask := new(task.CSDN)
	csdnTask.TaskName = "csdn"
	csdnTask.Init()
	csdnTask.Run(csdnTask.Worker,url.QueryEscape("区块链"))
}