package logger

import (
	"log"
	"fmt"
	"path/filepath"
	"runtime"
)
var (
	DefaultCallerDepth = 2


	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Level int
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Debug(v ...interface{}){
	setPrefix(DEBUG)
	log.Println(v...)
}


func Info(v ...interface{}){
	setPrefix(INFO)
	log.Println(v...)
}

func Error(v ...interface{}){
	setPrefix(ERROR)
	log.Println(v...)
}

func setPrefix(level Level)  {
	_, file, line, _ := runtime.Caller(DefaultCallerDepth)
	logPrefix := fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	log.SetPrefix(logPrefix)

}