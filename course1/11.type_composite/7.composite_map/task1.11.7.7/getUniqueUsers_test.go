package main

import "testing"

func testEq(a, b []User) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetUniqeUsres(t *testing.T) {
	users := []User{
		{Nickname: "zxc", Age: 18, Email: "WWW.y"},
		{Nickname: "ghoul", Age: 56, Email: "port"},
		{Nickname: "zxc", Age: 25, Email: "W.y"},
	}
	exp := []User{
		{Nickname: "zxc", Age: 18, Email: "WWW.y"},
		{Nickname: "ghoul", Age: 56, Email: "port"},
	}
	get := getUniqueUsers(users)
	if testEq(exp, get) == false {
		t.Errorf("want:%v, get:%v", exp, get)
	}
}
