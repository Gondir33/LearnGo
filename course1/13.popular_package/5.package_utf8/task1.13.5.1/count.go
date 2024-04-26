package main

func countUniqueUTF8Chars(s string) int {
	runes := []rune(s)
	res := make(map[rune]int, 0)
	for i := 0; i < len(runes); i++ {
		res[runes[i]]++
	}
	return len(res)
}
