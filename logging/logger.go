package logging

import (
	"io"
	"log"
	"os"
)

const (
	logPath = "logging/wirebo.log"
)

var (
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
)

func InitLoggers() {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	Error = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(mw, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
