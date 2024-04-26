package main

func RemoveIDX(xs []int, idx int) []int {
	if idx > len(xs) {
		return xs
	}
	return append(xs[:idx], xs[idx+1:]...)
}
