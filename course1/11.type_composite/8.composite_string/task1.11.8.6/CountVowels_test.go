package main

import "testing"

func TestCountVowels(t *testing.T) {
	input := []string{"Привет, мир!", "Hello, world!"}
	for i := 0; i < len(input); i++ {
		res := CountVowels(input[i])
		exp := 3
		if res != exp {
			t.Errorf("unlcuky")
		}
	}
}
