package main

import "testing"

func TestGetStringHeader(t *testing.T) {
	s := "Hello World!"
	res := getStringHeader(s)
	exp := getStringHeader(s)
	if res != exp {
		t.Errorf("want:%v, get:%v", exp, res)
	}
}
