package main

import (
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() int {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
	return c.value
}

func concurrentSafeCounter() int {
	counter := Counter{}
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	return counter.value
}
