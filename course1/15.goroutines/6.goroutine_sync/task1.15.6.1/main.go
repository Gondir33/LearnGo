package main

import (
	"fmt"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var wg sync.WaitGroup
	var result string
	var mutex sync.Mutex
	wg.Add(len(goroutines))
	for i := 0; i < len(goroutines); i++ {
		go func(n int) {
			defer wg.Done()
			mutex.Lock()
			result += goroutines[n]() + "\n"
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	return result
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %v", i)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}
