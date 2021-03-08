package zlog

import (
	"path"
	"runtime"
)

type LogLevel uint16

const (
	DEFAULT LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type Zlogger interface {
	Trace(format string, a ...interface{})

	Debug(format string, a ...interface{})

	Info(format string, a ...interface{})

	Warning(format string, a ...interface{})

	Error(format string, a ...interface{})

	Fatal(format string, a ...interface{})
}

func callinfo(skip int) (fileName, funcName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", "", line
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
