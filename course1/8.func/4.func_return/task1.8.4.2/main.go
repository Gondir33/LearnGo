package main

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) >= 1 {
		max := n[0]
		min := n[0]
		for i := 1; i < len(n); i++ {
			if max < n[i] {
				max = n[i]
			}
			if min > n[i] {
				min = n[i]
			}
		}
		return max, min
	} else {
		return 0, 0
	}
}
