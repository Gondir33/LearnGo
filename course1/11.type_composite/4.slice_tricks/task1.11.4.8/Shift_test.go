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
	res, res1 := Shift(xs)
	exp := []int{3, 1, 2}
	if testEq(res1, exp) == false || res != xs[0] {
		t.Errorf("input:%v, want:%v, get:%v", xs, exp, res)
	}
}
