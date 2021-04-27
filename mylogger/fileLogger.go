package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level      LogLevel
	filePath   string
	fileName   string
	fileObj    *os.File
	errFileObj *os.File
	maxSize    int64
}

type MsgInfo struct {
	level    string
	msg      string
	fileName string
	now      string
	line     int
}

var MsgChannel chan *MsgInfo = make(chan *MsgInfo, 50000)

func NewFileLogger(levStr, fp, fn string, ms int64) *FileLogger {
	ll, err := parse(levStr)
	if err != nil {
		panic("parse log lovel failed")
	}
	fo, err := os.OpenFile(path.Join(fp, fn), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open file failed")
	}
	errFo, err := os.OpenFile(path.Join(fp, fn)+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open err file failed")
	}
	fl := &FileLogger{
		Level:      ll,
		filePath:   fp,
		fileName:   fn,
		fileObj:    fo,
		errFileObj: errFo,
		maxSize:    ms,
	}
	go fl.runGoroutine()
	return fl
}

func (fl *FileLogger) runGoroutine() {
	for {
		m, ok := <-MsgChannel
		if ok {
			fl.fileObj = fl.checkSize(fl.fileObj)
			fmt.Fprintf(fl.fileObj, "[%s] [%s] [%s:%d]: %s\n", m.now, m.level, m.fileName, m.line, m.msg)
			l, err := parse(m.level)
			if err != nil {
				fmt.Println("unparse level failed")
				continue
			}
			if l >= ERROR {
				fl.errFileObj = fl.checkSize(fl.errFileObj)
				fmt.Fprintf(fl.errFileObj, "[%s] [%s] [%s:%d]: %s\n", m.now, m.level, m.fileName, m.line, m.msg)
			}
		}
	}
}

func (fl *FileLogger) fileLog(lv LogLevel, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	levelStr, err := unparse(lv)
	fileName, _, line := getInfo(3)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	MsgChannel <- &MsgInfo{
		level:    levelStr,
		msg:      msg,
		fileName: fileName,
		now:      now.Format("2006-01-02 15:04:05"),
		line:     line,
	}
}

func (fl *FileLogger) checkSize(file *os.File) *os.File {
	fileInfo, err := file.Stat()
	if err != nil {
		panic("get file info failed")
	}
	if fileInfo.Size() > fl.maxSize {
		nowStr := time.Now().Format("2006_0102_1504_05000")
		oldName := path.Join(fl.filePath, fileInfo.Name())
		newName := fmt.Sprintf("%s.bak.%s", oldName, nowStr)
		file.Close()
		os.Rename(oldName, newName)
		newFile, err := os.OpenFile(oldName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic("open new file failed")
		}
		return newFile
	}
	return file
}

// func (fl *FileLogger) checkDate(file *os.File) *os.File {
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		panic("get file info failed")
// 	}
// 	fileTime := fileInfo.ModTime()
// 	fileYear, fileMonth, fileDay := fileTime.Date()
// 	now := time.Now()
// 	year, month, day := now.Date()
// 	if year != fileYear || month != fileMonth || day != fileDay {
// 		nowStr := time.Now().Format("2006_0102_1504_05000")
// 		oldName := path.Join(fl.filePath, fileInfo.Name())
// 		newName := fmt.Sprintf("%s.bak.%s", oldName, nowStr)
// 		file.Close()
// 		os.Rename(oldName, newName)
// 		newFile, err := os.OpenFile(oldName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 		if err != nil {
// 			panic("open new file failed")
// 		}
// 		return newFile
// 	}
// 	return file
// }

func (fl *FileLogger) Debug(msg string, a ...interface{}) {
	if fl.Level <= DEBUG {
		fl.fileLog(DEBUG, msg, a...)
	}
}

func (fl *FileLogger) Trace(msg string, a ...interface{}) {
	if fl.Level <= TRACE {
		fl.fileLog(TRACE, msg, a...)
	}
}

func (fl *FileLogger) Info(msg string, a ...interface{}) {
	if fl.Level <= INFO {
		fl.fileLog(INFO, msg, a...)
	}
}

func (fl *FileLogger) Warning(msg string, a ...interface{}) {
	if fl.Level <= WARN {
		fl.fileLog(WARN, msg, a...)
	}
}

func (fl *FileLogger) Error(msg string, a ...interface{}) {
	if fl.Level <= ERROR {
		fl.fileLog(ERROR, msg, a...)
	}
}

func (fl *FileLogger) Fatal(msg string, a ...interface{}) {
	if fl.Level <= FATAL {
		fl.fileLog(FATAL, msg, a...)
	}
}

func (fl *FileLogger) Close() {
	fl.fileObj.Close()
	fl.errFileObj.Close()
}
