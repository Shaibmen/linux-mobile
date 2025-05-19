package logging

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}

type FileLogger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

var Log Logger

func InitLogger(path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	Log = &FileLogger{
		infoLog:  log.New(file, "[INFO]", log.LstdFlags),
		errorLog: log.New(file, "[ERROR]", log.LstdFlags),
	}
	return nil
}

func (f *FileLogger) Info(msg string) {
	f.infoLog.Println(msg)
}

func (f *FileLogger) Error(msg string, err error) {
	f.errorLog.Printf("%s: %v\n", msg, err)
}
