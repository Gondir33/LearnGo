package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		currentTime := time.Now()
		currentTimeFormatted := currentTime.Format("15:04:05")
		currentDateFormatted := currentTime.Format("2006-01-02")
		fmt.Printf("Текущее время: %s\n", currentTimeFormatted)
		fmt.Printf("Текущая дата: %s\n", currentDateFormatted)
		time.Sleep(1 * time.Second)
	}
}
