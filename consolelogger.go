package zlog

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	level LogLevel
}

func NewConsoleLogger(l LogLevel) *ConsoleLogger {
	return &ConsoleLogger{level: l}
}

func (c *ConsoleLogger) logEnable(level LogLevel) bool {
	return level >= c.level
}

func (c *ConsoleLogger) out(levelStr, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	t := time.Now()
	fileName, funcName, line := callinfo(3)
	fmt.Printf("[%s][%s:%s:%d][%s] %s", t.Format("2006/01/02 15:04:05"), fileName, funcName, line, levelStr, msg)
}

func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
	if c.logEnable(TRACE) {
		c.out("TRACE", format, a...)
	}
}

func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	if c.logEnable(DEBUG) {
		c.out("DEBUG", format, a...)
	}
}

func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	if c.logEnable(INFO) {
		c.out("INFO", format, a...)
	}
}

func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
	if c.logEnable(WARNING) {
		c.out("WARNING", format, a...)
	}
}

func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	if c.logEnable(ERROR) {
		c.out("ERROR", format, a...)
	}
}

func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	if c.logEnable(FATAL) {
		c.out("FATAL", format, a...)
	}
}
