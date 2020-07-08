package main

import (
	"github.com/itekton/gotesting/mylogger/customlogger"
)

var log customlogger.Logger

func consoleFile() {
	log = customlogger.NewConsoleLogger("info")
	log.Debug("debug 日志")
	log.Info("info日志")
}

func logFile() {
	log = customlogger.NewFileLogger("Info", "./", "access.log", 1<<30) // 文件日志实例 GB
	for {
		log.Debug("debug 日志")
		log.Info("info日志")
		log.Warning("这是一条warning日志")
		id := 10001
		name := "marion"
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
	}
}

func main() {
	//consoleFile()
	logFile()
}
