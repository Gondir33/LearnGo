package main

import "fmt"

func Sum(a ...int) int {
	var res int
	for _, arg := range a {
		res += arg
	}
	return res
}

func Mul(a ...int) int {
	res := 1
	for _, arg := range a {
		res *= arg
	}
	return res
}

func Sub(a ...int) int {
	var res int
	if len(a) > 0 {
		res = a[0]
		for i := 1; i < len(a); i++ {
			res -= a[i]
		}
	}
	return res
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))  // Output: 5
	fmt.Println(MathOperate(Mul, 1, 7, 3))  // Output: 21
	fmt.Println(MathOperate(Sub, 13, 2, 3)) // Output: 8
}
