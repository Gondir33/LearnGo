package workers

import (
	"fmt"
	"log"
	"os"
	"time"
)

const content = `---
menu:	
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: "/app/static/tasks/_index.md"

Должен меняться счетчик и время:

Текущее время: %v %v

Счетчик: %v



## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд`

func TimeWorker() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			currentTime := time.Now()
			err := os.WriteFile("/app/static/tasks/_index.md",
				[]byte(fmt.Sprintf(content, currentTime.Format("2006-01-02"), currentTime.Format("15:04:05"), b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
