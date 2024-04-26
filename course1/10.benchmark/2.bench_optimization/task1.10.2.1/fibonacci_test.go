package main

import (
	"math"
	"testing"
)

func FibonacciDP(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func FibonacciBinet(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}

// func GenerateTestData() []int {
// 	var data []int
// 	for i := 0; i < 100000000; i++ {
// 		data = append(data, rand.Int())
// 	}
// 	return data
// }

func BenchmarkFibonacciDP(b *testing.B) {
	// data := GenerateTestData()
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibonacciDP(10)
	}
}

func BenchmarkFibonacciBinet(b *testing.B) {
	// data := GenerateTestData()
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibonacciBinet(10)
	}
}
