package main

import (
	"reflect"
	"testing"
	"time"
)

type TestData struct {
	id      int
	options []OrderOption
	want    *Order
}

func TestNewOrder(t *testing.T) {
	tmp := time.Now()
	testCase := TestData{
		id: 1,
		options: []OrderOption{
			WithCustomerID("123"),
			WithItems([]string{"item1", "item2"}),
			WithOrderDate(tmp),
		},
		want: &Order{ID: 1, CustomerID: "123", Items: []string{"item1", "item2"}, OrderDate: tmp},
	}

	got := NewOrder(testCase.id, testCase.options...)

	if !reflect.DeepEqual(got, testCase.want) {
		t.Errorf("Got %+v, want %+v", got, testCase.want)
	}
}
