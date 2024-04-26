package main

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	res := make([]Animal, 3, 3)
	for i := 0; i < len(res); i++ {
		res[i].Type = gofakeit.AnimalType()
		res[i].Name = gofakeit.Name()
		res[i].Age = gofakeit.IntRange(18, 60)
	}
	return res
}

func preparePrint(animals []Animal) string {
	var res string
	for i := 0; i < len(animals); i++ {
		res += "Тип: " + animals[i].Type + ", Имя: " + animals[i].Name + ", Возраст: " + strconv.Itoa(animals[i].Age) + "\n"
	}
	return res
}
