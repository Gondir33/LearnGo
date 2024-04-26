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

func TestFilterDividers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := FilterDividers(input, 2)
	exp := []int{2, 4, 6, 8, 10}
	if testEq(res, exp) == false {
		t.Errorf("input:%v, want:%v, get:%v", input, exp, res)
	}

}
