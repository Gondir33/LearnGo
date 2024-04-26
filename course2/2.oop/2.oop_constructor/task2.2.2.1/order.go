package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(id string) OrderOption {
	return func(o *Order) {
		o.CustomerID = id
	}
}

func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items
	}
}

func WithOrderDate(time time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = time
	}
}

func NewOrder(id int, options ...OrderOption) *Order {
	o := &Order{ID: id}

	for _, option := range options {
		option(o)
	}
	return o
}

/*
func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}
*/
