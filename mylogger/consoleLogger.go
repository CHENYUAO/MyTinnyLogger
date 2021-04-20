package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel
}

func NewConsoleLogger(level string) *ConsoleLogger {
	l, err := parse(level)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		Level: l,
	}
}

func (cl *ConsoleLogger) consoleLog(lv LogLevel, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	levelStr, err := unparse(lv)
	fileName, _, line := getInfo(3)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Printf("[%s] [%s] [%s:%d]: %s\n", now.Format("2006-01-02 15:04:05"), levelStr, fileName, line, msg)
}

func (cl *ConsoleLogger) Debug(msg string, a ...interface{}) {
	if cl.Level <= DEBUG {
		cl.consoleLog(DEBUG, msg, a...)
	}
}

func (cl *ConsoleLogger) Trace(msg string, a ...interface{}) {
	if cl.Level <= TRACE {
		cl.consoleLog(TRACE, msg, a...)
	}
}

func (cl *ConsoleLogger) Info(msg string, a ...interface{}) {
	if cl.Level <= INFO {
		cl.consoleLog(INFO, msg, a...)
	}
}

func (cl *ConsoleLogger) Warning(msg string, a ...interface{}) {
	if cl.Level <= WARN {
		cl.consoleLog(WARN, msg, a...)
	}
}

func (cl *ConsoleLogger) Error(msg string, a ...interface{}) {
	if cl.Level <= ERROR {
		cl.consoleLog(ERROR, msg, a...)
	}
}

func (cl *ConsoleLogger) Fatal(msg string, a ...interface{}) {
	if cl.Level <= FATAL {
		cl.consoleLog(FATAL, msg, a...)
	}
}
