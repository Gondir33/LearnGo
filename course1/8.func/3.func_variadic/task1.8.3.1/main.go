package main

import (
	"strconv"
)

func UserInfo(name string, age int, cities ...string) string {
	var res string
	res = "Имя: " + name + ", возраст: " + strconv.Itoa(age) + ", города: "
	for _, city := range cities {
		res += city
		if city != cities[len(cities)-1] {
			res += ", "
		}
	}
	return res
}
