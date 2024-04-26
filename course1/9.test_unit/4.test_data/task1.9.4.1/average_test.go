package main

import (
	"math/rand"
	"testing"
)

func generateSlice(size int) []float64 {
	var testData []float64
	for i := 0; i < size; i++ {
		testData = append(testData, rand.Float64())
	}
	return testData
}

func TestAverage(t *testing.T) {
	testData := generateSlice(10)
	res := average(testData)
	expected := average(testData)
	if res != expected {
		t.Errorf("Unexpected result, Input: %v, Expected: %v, Get: %v", testData, expected, res)
	}
}
