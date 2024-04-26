package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	var res string
	tmp := strings.Split(sentence, " ")
	for i := 0; i < len(tmp); i++ {
		if filter[tmp[i]] == false {
			res += tmp[i] + " "
		}
	}
	return strings.TrimSuffix(res, " ")
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
