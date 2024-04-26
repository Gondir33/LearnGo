package main

import "strconv"

func UserInfo(name, city, phone string, age, weight int) string {
	return "Имя: " + name + ", Город: " + city + ", Телефон: " + phone + ", Возраст: " + strconv.Itoa(age) + ", Вес: " + strconv.Itoa(weight)
}
