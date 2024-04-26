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

func TestInsertAfterIDX(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	exp := []int{1, 2, 3, 6, 7, 8, 4, 5}
	res := InsertAfterIDX(xs, 2, 6, 7, 8)
	if testEq(res, exp) == false {
		t.Errorf("Input:%v, Want:%v, Get:%v", xs, exp, res)
	}
}
