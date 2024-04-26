package main

import "testing"

func TestGetUsers(t *testing.T) {
	res := getUsers()
	for i := 0; i < len(res); i++ {
		if res[i].Age < 18 && res[i].Age > 60 {
			t.Errorf("Unexpected  Age not between 18 and 60")
		}
		if res[i].Name == "" {
			t.Errorf("No name")
		}
	}
}

func TestPreparePrint(t *testing.T) {
	input := getUsers()
	res := preparePrint(input)
	exp := preparePrint(input)

	if res != exp {
		t.Errorf("Unexpected error: want:%v, Get:%v", exp, res)
	}
}
