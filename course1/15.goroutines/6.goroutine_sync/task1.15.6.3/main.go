package main

import "sync/atomic"

// Пример структуры счетчика
type Counter struct {
	count int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment() {
	// Ваш код для увеличения значения счетчика на 1
	atomic.AddInt64(&c.count, 1)
}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	// Ваш код для получения текущего значения счетчика
	return atomic.LoadInt64(&c.count)
}
