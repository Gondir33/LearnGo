package main

import "testing"

func TestCountBytes(t *testing.T) {

	res := countBytes("Привет, мир!")
	exp := 21
	if res != exp {
		t.Errorf("want:%v, get:%v", exp, res)
	}
}

func TestCountSymbols(t *testing.T) {
	res := countSymbols("Привет, мир!")
	exp := 12
	if res != exp {
		t.Errorf("want:%v, get:%v", exp, res)
	}
}
