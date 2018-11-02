package manager

import (
	"newsCaptorsTask/task"
	"fmt"
)

func init(){
	fmt.Println("task init")
	csdnTask := new(task.CSDN)
	csdnTask.TaskName = "csdn"
	csdnTask.Init()
	csdnTask.Run(csdnTask.Worker)
}