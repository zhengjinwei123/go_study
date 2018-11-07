package main

import (
	"time"
	"newsCaptorsTask/log"
	_ "newsCaptorsTask/manager"
)


func main() {
	log := log.MyLogger("captor")
	log.Debug("task server run.")
	time.Sleep(time.Hour * 24 * 360 * 100)
}
