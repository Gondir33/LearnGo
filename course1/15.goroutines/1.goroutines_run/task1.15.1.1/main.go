package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(5 * time.Second)

	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}
	/*
	   Результат работы программы:
	   Горутина завершила работу
	   Горутина завершила работу раньше, чем таймер сработал
	*/
}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		select {
		case <-timer.C:
			ch <- "Таймер сработал"
		case <-stop:
			ch <- "Горутина завершила работу раньше, чем таймер сработал"
		}
		close(ch)
	}()
	return ch
}
