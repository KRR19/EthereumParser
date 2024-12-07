package logger

import "fmt"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Error(msg string) {
	fmt.Println(msg + Red)
}
