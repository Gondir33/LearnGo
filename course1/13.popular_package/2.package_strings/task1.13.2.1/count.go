package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	lowTxt := strings.ToLower(txt)
	wordsTxt := strings.Fields(lowTxt)
	ans := make(map[string]int)

	for i := 0; i < len(wordsTxt); i++ {
		for j := 0; j < len(words); j++ {
			if words[j] == wordsTxt[i] {
				ans[words[j]]++
			}
		}
	}
	return ans
}

func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. 
        Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}

	result := CountWordsInText(txt, words)

	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
}
