package main

import "fmt"

func PrintNumbers(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
