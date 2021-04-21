package main

import (
	"MyTinnyLogger/mylogger"
	"time"
)

//var log mylogger.Logger

func main() {
	log := mylogger.NewFileLogger("DEBUG", "./", "test.log", 10*1024)
	for {
		//log = mylogger.NewConsoleLogger("DEBUG")
		log.Debug("this is a DEBUG message")
		log.Trace("this is a TARCE message")
		log.Warning("this is a WARN message")
		log.Error("this is a ERROR message")
		log.Fatal("this is a FATAL message")
		time.Sleep(time.Millisecond * 100)
		//log.Close()
	}
}
