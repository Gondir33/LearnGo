package main

func average(xs []float64) float64 {
	var sum float64
	for _, arg := range xs {
		sum += arg
	}
	return sum / float64(len(xs))
}
