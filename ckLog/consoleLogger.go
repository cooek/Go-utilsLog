package ckLog

import (
	"day05/utils"
	"fmt"
	"os"
	"time"
)

type ConsloeLogger struct {
	level Level
}

func NewConsoloLogger(level Level) *ConsloeLogger {
	logger := &ConsloeLogger{
		level: level,
	}
	return logger
}
func (c *ConsloeLogger) FLog(fomart string, args ...interface{}) {
	if c.level > DebugLevel {
		return
	}
	//拼接文件
	massage := fmt.Sprintf(fomart, args...)
	fileName, funcName, line := utils.GetCallerInfo(3)
	timenow := time.Now().Format("2006:01:02 15:04:05")
	logmsg := fmt.Sprintf("[%s][%s:%s][%d][%s]%s", timenow, fileName, funcName, line, GetStringModo(c.level), massage)
	fmt.Fprintln(os.Stdout, logmsg)
}
func (c *ConsloeLogger) Debug(fomart string, args ...interface{}) {
	c.FLog(fomart, args...)
}
func (c *ConsloeLogger) Info(fomart string, args ...interface{}) {
	c.FLog(fomart, args)
}
func (c *ConsloeLogger) Warn(fomart string, args ...interface{}) {
	c.FLog(fomart, args)
}
func (c *ConsloeLogger) Error(fomart string, args ...interface{}) {
	c.FLog(fomart, args)
}
