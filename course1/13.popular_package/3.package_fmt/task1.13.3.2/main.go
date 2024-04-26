package main

import "fmt"

// Пример использования функции
func main() {
	var num int = 10
	var str string = "Hello"

	fmt.Println(getVariableType(num)) // Вывод: "int"
	fmt.Println(getVariableType(str)) // Вывод: "string"
}

// Функция для получения типа переменной
func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
