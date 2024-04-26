package main

import "testing"

func TestConcatStrings(t *testing.T) {
	res := concatStrings("asd", "asd")
	exp := "asdasd"
	if exp != res {
		t.Errorf("lol")
	}
}
