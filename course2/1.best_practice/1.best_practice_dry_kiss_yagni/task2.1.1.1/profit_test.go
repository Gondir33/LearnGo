package main

import (
	"testing"
)

func TestSetProduct(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	if s.product != &testData {
		t.Errorf("SetProduct don't work expected: %v, get: %v", testData, *s.product)
	}
}
func TestGetAverageProfit(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	res := s.GetAverageProfit()
	expected := 3.00
	if res != expected {
		t.Errorf("Incorrect function work, expected: %v, get: %v", expected, res)
	}
}
func TestGetAverageProfitPercent(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	res := s.GetAverageProfitPercent()
	expected := 11.11111111111111
	if res != expected {
		t.Errorf("Incorrect function work, expected: %v, get: %v", expected, res)
	}
}
func TestGetCurrentProfit(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	res := s.GetCurrentProfit()
	expected := 15.00
	if res != expected {
		t.Errorf("Incorrect function work, expected: %v, get: %v", expected, res)
	}
}
func TestGetDifferenceProfit(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	res := s.GetDifferenceProfit()
	expected := 0.00
	if res != expected {
		t.Errorf("Incorrect function work, expected: %v, get: %v", expected, res)
	}
}

func cmpFloatSlices(v1, v2 []float64) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

func TestGetAllData(t *testing.T) {
	testData := Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{28, 29, 30, 31, 32},
		Buys:          []float64{25, 26, 27, 28, 29},
		CurrentPrice:  30,
		ProfitPercent: 14.56,
	}
	var s StatisticProfit
	s.SetProduct(&testData)
	res := s.getAllData()
	expected := []float64{3, 11.11111111111111, 15, 0}
	if !cmpFloatSlices(res, expected) {
		t.Errorf("Incorrect function work, expected: %v, get: %v", expected, res)
	}
}
func TestAverage(t *testing.T) {
	testData := []float64{1.5, 1.6, 1.7, 1.8}
	var s StatisticProfit
	res := s.Average(testData)
	expected := 1.65
	if res != expected {
		t.Errorf("Sum don't work expected: %v, get: %v", expected, res)
	}
}
func TestSum(t *testing.T) {
	testData := []float64{1.5, 1.6, 1.7, 1.8}
	var s StatisticProfit
	res := s.Sum(testData)
	expected := 6.6
	if res != expected {
		t.Errorf("Sum don't work expected: %v, get: %v", expected, res)
	}
}
