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

func TestCut(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	exp := []int{2, 3, 4}
	res := Cut(xs, 1, 3)
	if testEq(res, exp) == false {
		t.Errorf("Input:%v, Want:%v, Get:%v", xs, exp, res)
	}
}
