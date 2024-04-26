package main

import "testing"

func TestMaxDifference(t *testing.T) {
	data := []int{1, 6, 10}
	res := MaxDifference(data)
	expected := 9
	if res != expected {
		t.Errorf("input:%v, want:%v, get:%v", data, expected, res)
	}
}
