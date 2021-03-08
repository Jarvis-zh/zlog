package zlog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileLogger struct {
	level    LogLevel
	fileName string
	filePath string
	fp       *os.File
	maxSize  int64
}

func NewFileLogger(l LogLevel, filePath, fileName string, maxSize int64) *FileLogger {
	fullpath := filepath.Join(filePath, fileName)
	fp, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		level:    l,
		fileName: fileName,
		filePath: filePath,
		fp:       fp,
		maxSize:  maxSize,
	}
}

func (f *FileLogger) logEnable(level LogLevel) bool {
	return level >= f.level
}

func (f *FileLogger) isNeedSplit() bool {
	fileInfo, err := f.fp.Stat()
	if err != nil {
		fmt.Println("1", err)
		return false
	}
	curSize := fileInfo.Size()
	return curSize >= f.maxSize
}

func (f *FileLogger) splitFile() error {
	//1.close old log file
	f.fp.Close()
	//2. rename old log file to bak file
	oldFullPath := filepath.Join(f.filePath, f.fileName)
	timeStr := time.Now().Format("20060102150405000")
	bakFullPath := oldFullPath + ".bak" + timeStr
	err := os.Rename(oldFullPath, bakFullPath)
	if err != nil {
		fmt.Println("2", err)
		return err
	}
	//3. open new log file
	file, err := os.OpenFile(oldFullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("3", err)
		return err
	}
	f.fp = file
	return nil
}

func (f *FileLogger) out(levelStr, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	t := time.Now()
	fileName, funcName, line := callinfo(3)
	if f.isNeedSplit() {
		f.splitFile()
	}
	fmt.Fprintf(f.fp, "[%s][%s:%s:%d][%s] %s", t.Format("2006/01/02 15:04:05"), fileName, funcName, line, levelStr, msg)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	if f.logEnable(TRACE) {
		f.out("TRACE", format, a...)
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.logEnable(DEBUG) {
		f.out("DEBUG", format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.logEnable(INFO) {
		f.out("INFO", format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.logEnable(WARNING) {
		f.out("WARNING", format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.logEnable(ERROR) {
		f.out("ERROR", format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.logEnable(FATAL) {
		f.out("FATAL", format, a...)
	}
}

func (f *FileLogger) Close() {
	if f.fp != nil {
		f.fp.Close()
	}
}
