package main

import (
	"fmt"
	"strings"
)

func countWordOccurrences(text string) map[string]int {
	keys := strings.Split(text, " ")
	myMap := map[string]int{}
	for i := 0; i < len(keys); i++ {
		myMap[keys[i]]++
	}
	return myMap
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
