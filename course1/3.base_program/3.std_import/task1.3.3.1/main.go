package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return (math.Sqrt(x))
}

func main() {
	result := Sqrt(4)
	fmt.Println(result)
}
