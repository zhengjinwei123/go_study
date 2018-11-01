package main

import (
	"newsCaptorsTask/util/fileParser"
	"fmt"
	"runtime"
	"path/filepath"
	"os"
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
	p,err := util.GetFilePath("./conf/app.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	fmt.Println(os.Getwd())
	conf := fileParser.SetConfig(GetCurDir() + "/conf/app.ini")

	fmt.Println(conf.GetValue("mysql","port"))
	fmt.Println(conf.GetValue("base","appName"))
	fmt.Println(conf.GetValue("mysql","host"))
}
