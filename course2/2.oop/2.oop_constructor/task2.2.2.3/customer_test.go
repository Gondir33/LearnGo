package main

import (
	"reflect"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	data := NewCustomer(1,
		WithAccount(&SavingsAccount{}),
		WithName("abdul"),
	)
	want := &Customer{ID: 1,
		Account: &SavingsAccount{},
		Name:    "abdul",
	}
	if !reflect.DeepEqual(data, want) {
		t.Errorf("Want %+v, Get %+v", want, data)
	}
}
