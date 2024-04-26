package main

func sum(xs [8]int) int {
	var res int
	for i := 0; i < len(xs); i++ {
		res += xs[i]
	}
	return res
}

func average(xs [8]int) float64 {
	return float64(sum(xs)) / float64(len(xs))
}

func averageFloat(ys [8]float64) float64 {
	var sum float64
	for i := 0; i < len(ys); i++ {
		sum += ys[i]
	}
	return sum / float64(len(ys))
}

func reverse(xs [8]int) [8]int {
	var tmp int
	for i := 0; i < len(xs)/2; i++ {
		tmp = xs[i]
		xs[i] = xs[len(xs)-i-1]
		xs[len(xs)-i-1] = tmp
	}
	return xs
}
