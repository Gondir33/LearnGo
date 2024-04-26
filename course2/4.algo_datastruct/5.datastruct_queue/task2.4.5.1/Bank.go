package main

import (
	"fmt"

	"github.com/gomarkdown/markdown"
)

type Bank struct {
	queue []string
}

func (b *Bank) AddClient(client string) {
	if b.queue == nil {
		b.queue = make([]string, 0)
	}
	md := []byte(client)
	clientMd := markdown.ToHTML(md, nil, nil)
	b.queue = append(b.queue, string(clientMd))
}

func (b *Bank) ServeNextClient() string {
	if len(b.queue) == 0 {
		md := []byte("No clients in the queue")
		clientMd := markdown.ToHTML(md, nil, nil)
		return string(clientMd)
	}
	var res string
	res, b.queue = b.queue[0], b.queue[1:]
	return res
}

func main() {
	bank := Bank{}

	bank.AddClient("Client 1")
	bank.AddClient("Client 2")
	bank.AddClient("Client 3")

	fmt.Println(bank.ServeNextClient()) // Output: Client 1
	fmt.Println(bank.ServeNextClient()) // Output: Client 2
	fmt.Println(bank.ServeNextClient()) // Output: Client 3
	fmt.Println(bank.ServeNextClient()) // Output: No clients in the queue
}
