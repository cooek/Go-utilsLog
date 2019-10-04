package ckLog

import (
	"fmt"
	"os"
	path2 "path"
	"time"
	"utilsLog/utils"
)

type LogFileLogger struct {
	level     Level
	fileName  string
	filePath  string
	file      *os.File
	fileerror *os.File
	maxSize   int64
}
type Logger interface {
	Debug(level Level, fomart string, args ...interface{})
	Info(level Level, fomart string, args ...interface{})
	Warning(level Level, fomart string, args ...interface{})
	Fatal(level Level, fomart string, args ...interface{})
	Close()
}

//对外访问一定要净身出户 ，不要带别名！！！！
func LogNewFileLogger(filePath, fileName string) *LogFileLogger {

	Logobj := &LogFileLogger{
		filePath: filePath,
		fileName: fileName,
		maxSize:  1024 * 1024,
	}
	Logobj.init()
	return Logobj
}
func (l *LogFileLogger) init() {

	//path := fmt.Sprintf("%s%s", l.fileName, l.filePath)
	path := path2.Join(l.filePath, l.fileName+".log")
	fileobj, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open %s is not pass", path))
	}
	l.file = fileobj

	errname := fmt.Sprintf("%s.err", l.fileName)
	errfileobj, err := os.OpenFile(errname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("open pass is not create %s", err))
	}
	l.fileerror = errfileobj
}

func (l *LogFileLogger) log(level Level, fomart string, args ...interface{}) {

	if l.level > level {
		return
	}
	massage := fmt.Sprintf(fomart, args...)
	timenow := time.Now().Format("2006:01:02 15:04:05")
	fileName, funcName, line := utils.GetCallerInfo(3)
	logmsg := fmt.Sprintf("[%s][%s:%s][%d][%s]%s", timenow, fileName, funcName, line, GetStringModo(level), massage)
	if l.checkSplit(l.file) {
		fmt.Println("满足！！！！！！！！！")
		l.file = l.splitLogFile(l.file)
	}

	fmt.Fprintln(l.file, logmsg)
	if level >= ErrorLevel {
		//error级别的也检查拆分
		if l.checkSplit(l.fileerror) {
			fmt.Println("满足！！！！！！！！！")
			l.fileerror = l.splitLogFile(l.fileerror)
		}
		fmt.Fprintln(l.fileerror, logmsg)
	}

}

//检查文件是否超过限额
func (l *LogFileLogger) checkSplit(file *os.File) bool {
	filesinfo, _ := file.Stat()
	filesize := filesinfo.Size()
	return filesize >= l.maxSize
}

//封装一个切分日志文件的方法
func (l *LogFileLogger) splitLogFile(files *os.File) *os.File {
	//检查文件大小是否超过maxSize
	//切分文件 关闭 备份 新建
	name := files.Name()
	newupName := fmt.Sprintf("%s_%v.back", name, time.Now().Unix())
	files.Close()
	os.Rename(name, newupName)
	fileobj, err := os.OpenFile(newupName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开文件失败", err))
	}
	return fileobj
}

func (l *LogFileLogger) Debug(level Level, fomart string, args ...interface{}) {
	l.log(level, fomart, args...)
}
func (l *LogFileLogger) Info(level Level, fomart string, args ...interface{}) {
	l.log(level, fomart, args...)
}
func (l *LogFileLogger) Warning(level Level, fomart string, args ...interface{}) {
	l.log(level, fomart, args...)
}
func (l *LogFileLogger) Fatal(level Level, fomart string, args ...interface{}) {
	l.log(level, fomart, args...)
}
func (l *LogFileLogger) Close() {
	l.file.Close()
	l.fileerror.Close()
}
