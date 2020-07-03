package main

import "github.com/itekton/gotesting/mylogger/customlogger"

var log customlogger.Logger

func main() {
	log = customlogger.NewConsoleLogger("info")
	log.Debug("debug 日志")
	log.Info("info日志")
}
