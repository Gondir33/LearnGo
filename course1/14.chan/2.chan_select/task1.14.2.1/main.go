package main

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}
