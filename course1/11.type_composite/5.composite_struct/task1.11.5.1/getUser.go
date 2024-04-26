package main

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	res := make([]User, 10, 20)
	for i := 0; i < len(res); i++ {
		res[i].Name = gofakeit.Name()
		res[i].Age = gofakeit.IntRange(18, 60)
	}
	return res
}
func preparePrint(arg []User) string {
	var res string
	for i := 0; i < len(arg); i++ {
		res += "Имя: " + arg[i].Name + ", Возраст: " + strconv.Itoa(arg[i].Age) + "\n"
	}
	return res
}
