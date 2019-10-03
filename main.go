package main

import "day05/ckLog"

func main() {

	logger := ckLog.LogNewFileLogger("./", "log")
	logger.Debug(ckLog.DebugLevel, "测试哈哈")

	defer logger.Close()

}
