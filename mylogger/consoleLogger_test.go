package mylogger

import "testing"

func TestConsoleRun(t *testing.T) {
	log := NewConsoleLogger("TRACE")
	log.Debug("this is a debug test")
	log.Error("this is an error test")
	log.Fatal("this is a fatal test")
	log.Info("this is an info test")
	log.Trace("this is a trace test")
}
