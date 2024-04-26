package main

import (
	"fmt"
	"os"
	"path"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	err := os.MkdirAll(path.Dir(filePath), perm)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, perm)
}

func main() {
	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		fmt.Println("ошибка")
	}
}
