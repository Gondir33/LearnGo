package main

import (
	"math"
	"strconv"
)

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	a, err1 := strconv.ParseFloat(initialValue, 64)
	if err1 != nil {
		return 0, err1
	}
	b, err2 := strconv.ParseFloat(finalValue, 64)
	if err2 != nil {
		return 0, err2
	}
	if a == 0 {
		return 0, nil
	}
	return math.Round((b-a)/a*10000) / 100, nil
}
