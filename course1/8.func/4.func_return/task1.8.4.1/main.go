package main

func DivideAndRemainder(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}
