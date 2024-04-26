package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func monitorGoroutines(prevGoroutines int) {
	ticker := time.NewTicker(time.Millisecond * 300)
	NumGoroutine := runtime.NumGoroutine()
	for {
		select {
		case <-ticker.C:
			log.Println("Текущее количество горутин: ", NumGoroutine)
			prevGoroutines = runtime.NumGoroutine()
		default:
			percentage := float64(NumGoroutine) / float64(prevGoroutines)
			if percentage > 1.2 {
				log.Println("⚠️ Предупреждение: Количество горутин увеличилось более чем на 20%!")
			} else if percentage < 0.8 {
				log.Println("⚠️ Предупреждение: Количество горутин уменьшилось более чем на 20%!")
			}
			//wtf how many time I should wait btw
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func main() {
	g, _ := errgroup.WithContext(context.Background())

	// Мониторинг горутин
	go func() {
		monitorGoroutines(runtime.NumGoroutine())
	}()

	// Имитация активной работы приложения с созданием горутин
	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}

	// Ожидание завершения всех горутин
	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
