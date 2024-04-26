package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}
func (order *Order) RemoveDish(dish Dish) {
	if len(order.Dishes) == 1 {
		order.Dishes = make([]Dish, 0, 0)
	} else {
		for i := 0; i < len(order.Dishes); i++ {
			if order.Dishes[i] == dish {
				order.Dishes = append(order.Dishes[:i], order.Dishes[i+1:]...)
				break
			}
		}
	}
}
func (order *Order) CalculateTotal() {
	var res float64
	for i := 0; i < len(order.Dishes); i++ {
		res += order.Dishes[i].Price
	}
	order.Total = res
}

func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)
}
