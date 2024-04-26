package main

import "testing"

func TestGetUniqueWords(t *testing.T) {
	res := getUniqueWords("bar bar bar foo foo baz")
	exp := "bar foo baz"
	if res != exp {
		t.Errorf("want:%v, get %v", exp, res)
	}
}
