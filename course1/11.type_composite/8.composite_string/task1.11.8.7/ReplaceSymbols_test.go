package main

import "testing"

func TestReplaceSymbols(t *testing.T) {
	res := ReplaceSymbols("Hello, world!", 'o', '0')
	exp := "Hell0, w0rld!"
	if res != exp {
		t.Errorf("all my homies hate to make it ")
	}
}
