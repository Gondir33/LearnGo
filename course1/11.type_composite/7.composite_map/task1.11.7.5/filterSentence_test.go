package main

import "testing"

func TestFilter(t *testing.T) {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}
	filteredSentence := filterSentence(sentence, filter)
	exp := "Lorem dolor sit amet consectetur adipiscing"
	if exp != filteredSentence {
		t.Errorf("UNexpected error")
	}

}
