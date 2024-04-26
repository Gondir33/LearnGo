package main

import "testing"

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAppendInt(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	res := appendInt(input[0], input[1]...)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if testEq(res, expected) == false {
		t.Errorf("Input:%v, %v, Want:%v, Get:%v", input[0], input[1], expected, res)
	}
}
