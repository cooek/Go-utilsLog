package main

import "utilsLog/ckLog"

func main() {

	logger := ckLog.LogNewFileLogger("./", "log")
	logger.Debug(ckLog.DebugLevel, "测试哈哈")

	defer logger.Close()

}
