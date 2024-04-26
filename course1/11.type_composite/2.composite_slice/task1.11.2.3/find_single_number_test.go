package main

import "testing"

func TestFindSingleNumber(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	res := findSingleNumber(numbers)
	expected := 5
	if res != expected {
		t.Errorf("Input:%v, Want:%v, Get:%v", numbers, expected, res)
	}
}

func TestBitwiseXOR(t *testing.T) {
	input := []int{1, 3}
	res := bitwiseXOR(input[0], input[1])
	expected := 1 ^ 3
	if res != expected {
		t.Errorf("Input:%v, Want:%v, Get:%v", input, expected, res)
	}
}
