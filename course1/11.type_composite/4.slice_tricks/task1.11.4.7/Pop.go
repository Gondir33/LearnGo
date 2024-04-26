package main

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, xs
	}
	return xs[0], xs[1:]
}
