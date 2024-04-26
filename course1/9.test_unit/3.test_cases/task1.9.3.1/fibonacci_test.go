package main

import "testing"

type testData struct {
	a        int
	expected int
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func TestFibonacci(t *testing.T) {
	testCases := []testData{
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, tc := range testCases {
		result := Fibonacci(tc.a)
		if tc.expected != result {
			t.Errorf("Unexpected result. Input: %d, Expected: %d, Got: %d", tc.a, tc.expected, result)
		}
	}
}
