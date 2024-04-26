package main

import (
	"fmt"
)

type BrowserHistory struct {
	stack []string
}

func (h *BrowserHistory) Visit(url string) {
	if h.stack == nil {
		h.stack = make([]string, 0)
	}
	fmt.Printf("Посещение [%v]\n", url)
	h.stack = append(h.stack, url)
}

func (h *BrowserHistory) Back() {
	if h.stack == nil {
		return
	}
	if len(h.stack) == 0 {
		fmt.Println("Нет больше истории для возврата")
		return
	}
	fmt.Printf("Возврат к [%v]\n", h.stack[len(h.stack)-1])
	h.stack = h.stack[:len(h.stack)-1]
}

func (h *BrowserHistory) PrintHistory() {
	for i := len(h.stack) - 1; i >= 0; i-- {
		fmt.Println(h.stack[i])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.PrintHistory()
}
