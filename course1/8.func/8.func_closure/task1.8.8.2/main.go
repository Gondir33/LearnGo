package main

import "fmt"

func multiplier(factor float64) func(float64) float64 {
	return func(y float64) float64 {
		return factor * y
	}
}

func main() {
	// Пример использования функции multiplier
	m := multiplier(2.5)
	result := m(10)
	fmt.Println(result) // Вывод: 25
}
