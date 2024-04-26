package main

import "testing"

func Test_getType(t *testing.T) {
	var i interface{} = 42
	s := getType(i)
	exp := "int"
	if exp != s {
		t.Errorf("want:%v get:%v", exp, s)
	}
}
