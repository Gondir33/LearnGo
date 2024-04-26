package main

import "testing"

func TestSum(t *testing.T) {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	res := sum(xs)
	expected := 36
	if expected != res {
		t.Errorf("Unexpected result: Input:%v, Want:%v, Get:%v", xs, expected, res)
	}
}

func TestAverage(t *testing.T) {
	xs := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	res := average(xs)
	expected := 4.5
	if expected != res {
		t.Errorf("Unexpected result: Input:%v, Want:%v, Get:%v", xs, expected, res)
	}
}

func TestAverageFloat(t *testing.T) {
	ys := [...]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}
	res := averageFloat(ys)
	expected := float64(5)
	if expected != res {
		t.Errorf("Unexpected result: Input:%v, Want:%v, Get:%v", ys, expected, res)
	}
}

func TestReverse(t *testing.T) {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	res := reverse(xs)
	expected := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	if res != expected {
		t.Errorf("Unexpected result: Input:%v, Want:%v, Get:%v", xs, expected, res)
	}
}
