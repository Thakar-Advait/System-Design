package logger

import (
	"fmt"
	"sync"
)

type Logger struct{}

var logger *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
	once.Do(func() {
		fmt.Println("Creating new logger instance...")
		logger = &Logger{}
	})
	return logger
}

func (l *Logger) Log(message string) string {
	return fmt.Sprintf("[LOG]: %s", message)
}
