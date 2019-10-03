package utils

import (
	"path"
	"runtime"
)

func GetCallerInfo(skip int) (fileName, funcName string, line int) {
	//文件，行号，程序是否ok
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {
		return
	}
	//拿到最后的文件名
	fileName = path.Base(file)
	//拿到函数的名字
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //函数名最后一行
	return fileName, funcName, line
}
