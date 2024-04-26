package main

import "testing"

func TestCreateUniqueText(t *testing.T) {
	res := createUniqueText("bar bar bar foo foo baz")
	exp := "bar foo baz"
	if res != exp {
		t.Errorf("want:%v, get %v", exp, res)
	}
}
