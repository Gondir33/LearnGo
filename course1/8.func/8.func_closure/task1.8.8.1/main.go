package main

import "fmt"

func createCounter() func() int {
	var x int
	return func() int {
		x = x + 1
		return x
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
