package logging

import (
	"log"
	"os"
)

type FileLogger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewFileLogger(path string) (*FileLogger, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		infoLog:  log.New(file, "[INFO]", log.LstdFlags),
		errorLog: log.New(file, "[ERROR]", log.LstdFlags),
	}, nil
}

func (f *FileLogger) Info(msg string) {
	f.infoLog.Println(msg)
}

func (f *FileLogger) Error(msg string) {
	f.errorLog.Println(msg)
}
