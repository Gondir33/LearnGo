package main

import (
	"log"
	"time"
)

type Task struct {
	ID     int
	Data   string
	Status bool
}

type WorkerPool struct {
	Queue        chan *Task
	WorkersCount int
	Result       chan *Task
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		Queue:        make(chan *Task),
		WorkersCount: workers,
		Result:       make(chan *Task),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.WorkersCount; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	for {
		task := <-wp.Queue
		task.Status = true
		log.Printf("Executing task %d: %s\n", task.ID, task.Data)
		wp.Result <- task
	}
}

func (wp *WorkerPool) AddTask(task *Task) {
	wp.Queue <- task
}

// Wait - waiting for all tasks to complete in seconds
func (wp *WorkerPool) Wait(timer int) {
	delay := time.Duration(int(time.Second) * timer)
	<-time.After(delay)
}

func main() {
	pool := NewWorkerPool(5)
	pool.Start()

	tasks := []*Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
	}

	go func() {
		for _, task := range tasks {
			pool.AddTask(task)
		}
	}()

	go func() {
		for v := range pool.Result {
			log.Printf("Task %d completed: %s status: %v\n", v.ID, v.Data, v.Status)
		}
	}()

	// Waiting for all tasks to complete
	pool.Wait(1)
}

/*
2023/07/19 00:01:39 Executing task 3: Task 3
2023/07/19 00:01:39 Task 3 completed: Task 3 status: true
2023/07/19 00:01:39 Executing task 1: Task 1
2023/07/19 00:01:39 Task 1 completed: Task 1 status: true
2023/07/19 00:01:39 Executing task 2: Task 2
2023/07/19 00:01:39 Task 2 completed: Task 2 status: true
*/
