package main

func RemoveExtraMemory(xs []int) []int {
	res := make([]int, len(xs), len(xs))
	for i := 0; i < len(xs); i++ {
		res[i] = xs[i]
	}
	return res
}
