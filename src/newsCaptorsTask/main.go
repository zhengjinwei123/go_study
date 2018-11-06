package main

import (
	_ "newsCaptorsTask/manager"
	"time"
	"newsCaptorsTask/manager"
)

var log *manager.JadeLogger

func init(){
	log = new(manager.JadeLogger).SetType("captor").GetLogger()
}

func main() {
	log.Debug("task server run.")
	time.Sleep(time.Hour * 24 * 360 * 100)
}
