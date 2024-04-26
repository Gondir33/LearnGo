package main

import (
	"bytes"
	"fmt"
	"os"
)

func WriteFile(buf *bytes.Buffer, file *os.File) error {
	_, err := file.Write(buf.Bytes())
	return err
}

func main() {
	filePath := "/home/gondir/go-kata/course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"
	// Открываем файл для записи
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close() // отложенная функция закрытия дескриптора файла

	err = WriteFile(bytes.NewBufferString("Hello, World!"), file)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
	}
}
