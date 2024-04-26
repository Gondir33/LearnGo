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

func TestMergeMaps(t *testing.T) {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	mergedMap := mergeMaps(map1, map2)
	exp := map[string]int{
		"banana": 2,
		"orange": 5,
		"grape":  4,
		"apple":  3,
	}
	if cmpMap(exp, mergedMap) == false {
		t.Errorf("not Merged")
	}

}
