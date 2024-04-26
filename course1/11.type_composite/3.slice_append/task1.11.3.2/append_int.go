package main

func appendInt(xs *[]int, x ...int) {
	*xs = append(*xs, x...)
}
