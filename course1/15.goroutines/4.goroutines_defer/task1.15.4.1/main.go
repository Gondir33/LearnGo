package main

import (
	"os"
)

func writeToFile(file *os.File, data string) error {
	defer file.Close()
	_, err := file.WriteString(data)
	return err
}
