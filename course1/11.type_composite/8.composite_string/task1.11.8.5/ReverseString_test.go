package main

import "testing"

func TestReverseString(t *testing.T) {
	res := ReverseString("Hello, world!")
	exp := "!dlrow ,olleH"
	if res != exp {
		t.Errorf("Fuck tests all my homies use brain")
	}
}
