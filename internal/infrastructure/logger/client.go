package logger

import (
	"fmt"
	"io"
)

type Logger struct {
	output io.Writer
}

func NewLogger(output io.Writer) *Logger {
	return &Logger{output: output}
}

func (l *Logger) Info(msg string) {
	fmt.Fprintln(l.output, msg)
}

func (l *Logger) Error(msg string) {
	fmt.Fprintln(l.output, Red+msg+Reset)
}
