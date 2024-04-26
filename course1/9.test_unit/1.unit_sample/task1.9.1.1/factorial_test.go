package main

import "testing"

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return Factorial(n-1) * n
}

func TestFactorial(t *testing.T) {
	var result int
	result = Factorial(0)
	if result != 1 {
		t.Errorf("Factorial(0) = %d; want 1", result)
	}
	result = Factorial(1)
	if result != 1 {
		t.Errorf("Factorial(1) = %d; want 1", result)
	}
	result = Factorial(5)
	if result != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result)
	}
}
