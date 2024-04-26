// Код сервера
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func chatServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func main() {
	chatServer()
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool, 10)

	for {
		select {
		case mess := <-messages:
			for cli, life := range clients {
				if life == true {
					cli.conn.Write([]byte(mess))
				}
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			clients[cli] = false
		}
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		if len(input.Text()) == 0 {
			break
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	select {
	case mess := <-ch:
		conn.Write([]byte(mess))
	}
}
