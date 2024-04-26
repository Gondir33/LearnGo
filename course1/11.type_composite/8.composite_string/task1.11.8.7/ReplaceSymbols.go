package main

import "strings"

func ReplaceSymbols(s string, old, new rune) string {
	return strings.ReplaceAll(s, string(old), string(new))
}
