package mylogger

import (
	"testing"
	"time"
)

func TestFileRun(t *testing.T) {
	log := NewFileLogger("TRACE", "./", "test.log", 10*1024*1024)
	log.Debug("this is a debug test")
	log.Error("this is an error test")
	log.Fatal("this is a fatal test")
	log.Info("this is an info test")
	log.Trace("this is a trace test")
	time.Sleep(time.Second)
}
