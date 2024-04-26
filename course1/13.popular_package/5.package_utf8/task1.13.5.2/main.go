package main

import (
	"fmt"
	"strings"
)

func isRussianLetter(char rune) bool {
	chars := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	for i := 0; i < len(chars); i++ {
		if char == chars[i] {
			return true
		}
	}
	return false
}

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	s = strings.ToLower(s)
	for _, char := range s {
		if isRussianLetter(char) {

			counts[char]++
		}
	}

	return counts
}

func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2
	}
}
