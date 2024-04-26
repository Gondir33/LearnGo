package main

import (
	"time"
)

const (
	RECURSIVE = "recursive"
	ITERATIVE = "iterative"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return factorialRecursive(n-1) * n
}

func factorialIterative(n int) int {
	res := 1
	for ; n > 1; n-- {
		res *= n
	}
	return res
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() map[string]bool {
	res := make(map[string]bool)
	since := time.Now()
	factorialRecursive(100000)
	recursiveTime := time.Since(since)
	since = time.Now()
	factorialIterative(100000)
	iterativeTime := time.Since(since)
	res[RECURSIVE] = recursiveTime > iterativeTime
	res[ITERATIVE] = recursiveTime < iterativeTime
	return res
}

/*
func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}
*/
