package main

import (
	"newsCaptorsTask/util/fileParser"
	"fmt"
	"runtime"
	"path/filepath"
	"newsCaptorsTask/util"
)

func GetCurDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("11111")
	}
	return filepath.Dir(file)
}

func main() {
	_,err := util.GetFilePath("./conf/app.ini")
	if err != nil {
		panic(err)
	}
	conf := fileParser.SetConfig(GetCurDir() + "/conf/app.ini")

	fmt.Println(conf.GetValue("mysql","port"))
	fmt.Println(conf.GetValue("base","appName"))
	fmt.Println(conf.GetValue("mysql","host"))
}
