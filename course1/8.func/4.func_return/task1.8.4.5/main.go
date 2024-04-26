package main

import "math"

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	return math.Round((finalValue-initialValue)/initialValue*10000) / 100
}
