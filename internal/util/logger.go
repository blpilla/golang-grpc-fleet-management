// internal/util/logger.go
package util

import (
	"log"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.Default()}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Logger.SetPrefix("FATAL: ")
	l.Logger.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Logger.SetPrefix("INFO: ")
	l.Logger.Println(v...)
}
