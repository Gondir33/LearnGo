package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

type Logger interface {
	Log(string) error
}

type ConsoleLogger struct {
	Writer *os.File
}

func (c *ConsoleLogger) Log(mess string) error {
	_, err := c.Writer.Write([]byte(mess))
	return err
}

type FileLogger struct {
	File *os.File
}

func (f *FileLogger) Log(mess string) error {
	_, err := f.File.Write([]byte(mess))
	return err
}

type RemoteLogger struct {
	api string
}

func (r *RemoteLogger) Lof(mess string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", r.api, bytes.NewBuffer([]byte(mess)))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, _ := os.Create("txt")
	fileLogger := &FileLogger{File: file} // Здесь замени на открытие реального файла, но для примера мы будем использовать os.Stdout

	loggers := []Logger{consoleLogger, fileLogger}
	LogAll(loggers, "This is a test log message.")
}
