# console
将日志信息输入到屏幕
- `customlogger：` 主要逻辑代码目录
- `log_test：` 目录为测试代码实例
```go
package main

import "github.com/itekton/gotesting/mylogger/customlogger"

var log customlogger.Logger

func main() {
	log = customlogger.NewConsoleLogger("info")
	log.Debug("debug 日志")
	log.Info("info日志")
}
/*
输出信息为：
[2020-07-03 13:50:00] [INFO] [main.go:main:10] info日志
*/
```

# file
```go
func logFile() {
	log = mylogger.NewFileLogger("Info", "./", "access.log", 1<<30) // 文件日志实例 GB
	log.Debug("debug 日志")
	log.Info("info日志")
	log.Warning("这是一条warning日志")
	id := 10001
	name := "marion"
	log.Error("这是一条Error日志,id:%d,name:%s", id, name)
}
```