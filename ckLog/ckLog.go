package ckLog

//定义类型代表日志级别
type Level uint16

const (
	DebugLevel Level = iota
	InfoLevel
	WarngLevel
	ErrorLevel
	FatalLevel
)

func GetStringModo(level Level) string {

	switch level {
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarngLevel:
		return "Warng"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	default:
		return "Debug"
	}
}
