package main

import (
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}

func (f *FileLogger) Log(message string) {
	f.file.Write([]byte(message))
}

type LogSystem struct {
	fileLogger FileLogger
}

func (l *LogSystem) Log(message string) {
	l.fileLogger.Log(message)
}

// LogOption functional option type
type LogOption func(*LogSystem)

func WithLogger(file FileLogger) LogOption {
	return func(l *LogSystem) {
		l.fileLogger = file
	}
}

func NewLogSystem(options ...LogOption) *LogSystem {
	l := &LogSystem{}

	for _, option := range options {
		option(l)
	}
	return l
}

/*
func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}
*/
