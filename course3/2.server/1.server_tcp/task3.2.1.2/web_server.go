package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// Парсим HTTP-запрос
	method, path := parseRequest(request)

	// Готовим HTTP-ответ
	var response string
	if method == "GET" && path == "/" {
		response = "HTTP/1.1 200 OK\nContent-Type: text/html\n\n<!DOCTYPE html>\n<html>\n<head>\n<title>Webserver</title>\n</head>\n<body>\nhello world\n</body>\n</html>"
	} else {
		response = "HTTP/1.1 404 Not Found\n\n404 Not Found"
	}

	// Отправляем HTTP-ответ клиенту
	conn.Write([]byte(response))
}

func parseRequest(request string) (string, string) {
	parts := strings.Fields(request)
	if len(parts) >= 2 {
		return parts[0], parts[1]
	}
	return "", ""
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
