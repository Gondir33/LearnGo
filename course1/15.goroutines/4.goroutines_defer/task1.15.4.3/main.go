package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 1)
	myPanic(ch)
	fmt.Println(<-ch)
}

func myPanic(ch chan string) {
	defer close(ch)
	defer func() {
		if r := recover(); r != nil {
			ch <- fmt.Sprint(r)
		}
	}()
	panic("my panic message")
}
