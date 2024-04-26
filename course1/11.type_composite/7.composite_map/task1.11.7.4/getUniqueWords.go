package main

import "strings"

func getUniqueWords(text string) string {
	var res string
	tmp := strings.Split(text, " ")
	myMap := map[string]int{}
	for i := 0; i < len(tmp); i++ {
		myMap[tmp[i]]++
		if myMap[tmp[i]] == 1 {
			res += tmp[i] + " "
		}
	}

	return strings.TrimSuffix(res, " ")
}
