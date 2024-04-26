package main

import (
	"sync"
)

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(from))

	for _, ch := range from {
		go func(ch chan int) {
			defer wg.Done()
			for {
				val, ok := <-ch
				if !ok {
					return
				}
				mergeTo <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()
}

func mergeChan2(chans ...chan int) chan int {
	mergeTo := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(ch chan int) {
			defer wg.Done()
			for {
				val, ok := <-ch
				if !ok {
					return
				}
				mergeTo <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()

	return mergeTo
}
