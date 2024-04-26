package main

import "testing"

func cmpMap(a map[string]int, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}

func TestCountWordOccurrences(t *testing.T) {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)
	exp := map[string]int{
		"sit":         1,
		"amet":        1,
		"consectetur": 1,
		"adipiscing":  1,
		"elit":        1,
		"Lorem":       1,
		"ipsum":       2,
		"dolor":       1,
	}
	if cmpMap(occurrences, exp) == false {
		t.Errorf("want:%v, get:%v", exp, occurrences)
	}
}
