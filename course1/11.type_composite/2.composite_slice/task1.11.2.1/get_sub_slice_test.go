package main

import (
	"testing"
)

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

func TestGetSubSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := getSubSlice(numbers, 2, 6)
	expected := []int{3, 4, 5, 6}
	if testEq(res, expected) == false {
		t.Errorf("Unexpected result, Input:%v, Want:%v, Get:%v", numbers, expected, res)
	}

}
