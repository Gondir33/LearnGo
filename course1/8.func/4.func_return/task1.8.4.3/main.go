package main

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price * float64(quantity), price
}
