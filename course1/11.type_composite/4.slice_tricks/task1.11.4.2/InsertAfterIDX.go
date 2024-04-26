package main

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	tmp := make([]int, len(x), len(x))
	for i := 0; i < len(x); i++ {
		tmp[i] = x[i]
	}
	return append(xs[:idx+1], append(tmp, xs[idx+1:]...)...)
}
