package config

import (
	"newsCaptorsTask/util/fileParser"
	"newsCaptorsTask/util"
)

var C *fileParser.Config

func init() {
	p, err := util.GetFilePath("../conf/app.ini")
	if err != nil {
		panic(err)
	}
	C = fileParser.SetConfig(p)
}

func AppConf() *fileParser.Config {
	return C
}
