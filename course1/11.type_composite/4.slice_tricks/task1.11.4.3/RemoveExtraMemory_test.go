package main

import "testing"

func TestRemoveExtraMemory(t *testing.T) {
	input := make([]int, 5, 12)
	res := RemoveExtraMemory(input)
	exp := make([]int, 5, 5)
	if cap(res) != cap(exp) {
		t.Errorf("inpit cap:%v, want:%v, get:%v", cap(input), cap(exp), cap(res))
	}
}
