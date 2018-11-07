package log

import (
	"github.com/op/go-logging"
	"newsCaptorsTask/util"
	"newsCaptorsTask/config"
	"os"
	"fmt"
	"time"
)

var g_filePath string

var g_mapLogger map[string]*JadeLogger

func init(){
	g_filePath,_ = util.GetFilePath(config.AppConf().GetValue("log","output"))
	g_mapLogger = make(map[string]*JadeLogger,20)
}

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)


type JadeLogger struct {
	logType string
	curDateTimeStr string
	log *logging.Logger
}

func (this *JadeLogger) filePath(fileType string) string{
	dir := util.ToLocaleDateString()
	this.curDateTimeStr = dir

	exists,_ := util.PathExists(g_filePath + "\\" + dir)
	if !exists {
		err := os.Mkdir(g_filePath + "\\" + dir,os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	fileDir := g_filePath + "\\" + dir + "\\"+ fileType + ".log"
	return fileDir
}

func (this *JadeLogger) SetType (logType string) *JadeLogger{
	this.logType = logType

	tmpPath := this.filePath(logType)

	logFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Println("file path:",tmpPath)
		panic(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.DEBUG, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)

	return this
}

func (this *JadeLogger) getLogger() *JadeLogger{
	this.log = logging.MustGetLogger("example")
	return this
}

func (this *JadeLogger) prefix() string {
	return "[" + util.ToLocaleDateTimeString() +"] "
}

func (this *JadeLogger) Debug(args ...interface{}) {
	this.flushDir()
	this.log.Debug(this.prefix(),args)
}

func (this *JadeLogger) Warning(args ...interface{}) {
	this.flushDir()
	this.log.Warning(this.prefix(),args)
}

func (this *JadeLogger) Error(args ...interface{}) {
	this.flushDir()
	this.log.Error(this.prefix(),args)
}

func (this *JadeLogger) Info(args ...interface{}) {
	this.flushDir()
	this.log.Info(this.prefix(),args)
}

func (this *JadeLogger) flushDir(){
	now := time.Now()
	str := util.ToDateTimeString(now.Year(),int(now.Month()),now.Day(),0,0,0)
	diffSec := util.TimeSecDiff(str,this.curDateTimeStr+" 00:00:00")

	if diffSec == 0 {
		return
	}

	dir := util.ToLocaleDateString()
	exists,_ := util.PathExists(g_filePath + "\\" + dir)
	if !exists {
		err := os.Mkdir(g_filePath + "\\" + dir,os.ModePerm)
		if err != nil {
			panic(err)
		}
		this.SetType(this.logType)
	}
}


func MyLogger(logType string) *JadeLogger{
	_log,ok := g_mapLogger[logType]
	if !ok {
		g_mapLogger[logType] = new(JadeLogger).SetType(logType).getLogger()
		return g_mapLogger[logType]
	}
	return _log
}
