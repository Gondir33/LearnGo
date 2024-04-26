package main

import "fmt"

func calculate(a int, b int) (int, int, int, int, int) {
	return a + b, a - b, a * b, a / b, a % b
}

func main() {
	var a int
	var b int
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int // объявить явно все переменные c явным указанием типов, по отдельности
	a, b = 10, 3
	sum, difference, product, quotient, remainder = calculate(a, b)
	fmt.Printf("a = %d b = %d sum = %d difference = %d product = %d quotient = %d remainder = %d", a, b, sum, difference, product, quotient, remainder)
}
