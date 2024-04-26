package main

import "testing"

func TestAddDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	order.AddDish(dish1)
	if order.Dishes[0] != dish1 {
		t.Errorf("don't add the dish")
	}
}
func TestRemoveDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}
	order.AddDish(dish1)
	order.AddDish(dish2)
	order.RemoveDish(dish1)
	if order.Dishes[0] != dish2 {
		t.Errorf("don't remove the dish")
	}
}
func TestCalculateTotal(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	if order.Total != 16.98 {
		t.Errorf("don't calculate total")
	}
}
