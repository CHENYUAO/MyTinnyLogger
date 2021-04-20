package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type Logger interface {
	Debug(msg string, a ...interface{})
	Info(msg string, a ...interface{})
	Trace(msg string, a ...interface{})
	Warning(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
}

func parse(levelStr string) (LogLevel, error) {
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKNOW, errors.New("invalid log level")
	}
}

func unparse(l LogLevel) (string, error) {
	switch l {
	case TRACE:
		return "TRACE", nil
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARN:
		return "WARN", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		return "UNKNOWN", errors.New("invaild log level")
	}

}

func getInfo(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller()failed")
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)
	return fileName, funcName, line
}
