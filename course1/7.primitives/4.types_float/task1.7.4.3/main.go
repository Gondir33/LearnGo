package main

import "math"

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	aRound := math.Round(a*math.Pow10(decimalPlaces)) / math.Pow10(decimalPlaces)
	bRound := math.Round(b*math.Pow10(decimalPlaces)) / math.Pow10(decimalPlaces)
	isEqual = false
	difference = math.Abs(aRound - bRound)
	if aRound == bRound {
		isEqual = true
	}
	return isEqual, difference
}
