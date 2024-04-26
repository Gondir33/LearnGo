package main

import "strings"

const (
	Vowels = "aoeiuyаеёиоуыэюя"
)

func CountVowels(s string) int {
	var res int
	s = strings.ToLower(s)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if strings.ContainsRune(Vowels, runes[i]) {
			res++
		}
	}
	return res
}
