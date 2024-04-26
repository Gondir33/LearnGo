package main

func InsertToStart(xs []int, x ...int) []int {
	a := make([]int, len(x), len(x))
	for i := 0; i < len(x); i++ {
		a[i] = x[i]
	}
	return append(a, xs...)
}
