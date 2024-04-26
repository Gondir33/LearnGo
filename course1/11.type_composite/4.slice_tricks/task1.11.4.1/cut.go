package main

func Cut(xs []int, start, end int) []int {
	if start > len(xs) || end > len(xs) {
		return []int{}
	}
	return xs[start : end+1]
}
