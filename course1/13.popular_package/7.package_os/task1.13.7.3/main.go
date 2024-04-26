package main

import (
	"os"
)

func ReadString(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
