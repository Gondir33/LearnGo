// Код клиента
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func chatClient() {
	// подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// запустить горутину, которая будет читать все сообщения от сервера и выводить их в консоль
	go clientReader(conn)

	// читать сообщения от stdin и отправлять их на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write(scanner.Bytes())
		if len(scanner.Text()) == 0 {
			break
		}
	}
}

func main() {
	chatClient()
}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	buffer := make([]byte, 1000)

	for {
		end, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении данных:", err)
			break
		}
		fmt.Println(string(buffer[:end]))
	}
}
