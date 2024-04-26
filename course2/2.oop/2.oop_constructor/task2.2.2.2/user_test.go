package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	id      int
	options []UserOption
	want    *User
}

func TestNewOrder(t *testing.T) {
	testCase := TestData{
		id: 1,
		options: []UserOption{
			WithUsername("testuser"),
			WithEmail("testuser@example.com"),
			WithRole("admin"),
		},
		want: &User{ID: 1, Username: "testuser", Email: "testuser@example.com", Role: "admin"},
	}

	got := NewUser(testCase.id, testCase.options...)

	if !reflect.DeepEqual(got, testCase.want) {
		t.Errorf("Got %+v, want %+v", got, testCase.want)
	}
}
