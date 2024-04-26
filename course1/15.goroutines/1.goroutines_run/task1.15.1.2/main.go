package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	ch := make(chan string)

	go func() {
		select {
		case <-time.After(d + time.Duration(time.Millisecond*100)):
			ticker.Stop()
			close(ch)
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				ch <- message
			}
		}
	}()

	return ch
}
