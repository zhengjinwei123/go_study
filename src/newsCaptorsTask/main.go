package main

import (
	_ "newsCaptorsTask/manager"
	"time"
	"newsCaptorsTask/manager"
)


func main() {
	log := manager.MyLogger("captor")
	log.Debug("task server run.")
	time.Sleep(time.Hour * 24 * 360 * 100)
}
