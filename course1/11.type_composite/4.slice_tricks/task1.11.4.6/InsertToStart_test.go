package main

import "testing"

func testEq(a []int, b []int) bool {
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

func TestInsertToStart(t *testing.T) {
	xs := []int{1, 2, 3}
	res := InsertToStart(xs, 4, 5, 6)
	exp := []int{4, 5, 6, 1, 2, 3}
	if testEq(res, exp) == false {
		t.Errorf("input:%v, want:%v, get:%v", xs, exp, res)
	}
}
