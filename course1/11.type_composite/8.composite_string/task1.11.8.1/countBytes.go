package main

import "unicode/utf8"

func countBytes(s string) int {
	bytes := []byte(s)
	return len(bytes)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}
