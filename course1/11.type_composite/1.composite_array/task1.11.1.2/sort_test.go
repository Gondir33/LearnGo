package main

import "testing"

func TestSortDescInt(t *testing.T) {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	res := sortDescInt(xs)
	expected := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	if res != expected {
		t.Errorf("Unexpected result, Input:%v, Want:%v, Get%v", xs, expected, res)
	}
}

func TestSortAscInt(t *testing.T) {
	xs := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	res := sortAscInt(xs)
	expected := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	if res != expected {
		t.Errorf("Unexpected result, Input:%v, Want:%v, Get%v", xs, expected, res)
	}
}

func TestSortDescFloat(t *testing.T) {
	xs := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}
	res := sortDescFloat(xs)
	expected := [8]float64{8, 7, 6, 5, 4, 3, 2, 1}
	if res != expected {
		t.Errorf("Unexpected result, Input:%v, Want:%v, Get%v", xs, expected, res)
	}
}

func TestSortAscFloat(t *testing.T) {
	xs := [8]float64{8, 7, 6, 5, 4, 3, 2, 1}
	res := sortAscFloat(xs)
	expected := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}
	if res != expected {
		t.Errorf("Unexpected result, Input:%v, Want:%v, Get%v", xs, expected, res)
	}
}
