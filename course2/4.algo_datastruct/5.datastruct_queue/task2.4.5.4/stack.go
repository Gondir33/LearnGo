package main

import "fmt"

// Пример кода на языке программирования
var stack = []int{}

// Определение стека

// Функция для добавления элемента в стек
func push(value int) {
	stack = append(stack, value)
}

// Функция для удаления и возврата последнего элемента из стека
func pop() int {
	if len(stack) == 0 {
		return 0
	}
	var res int
	res, stack = stack[len(stack)-1], stack[:len(stack)-1]
	return res
}

// Пример использования стека для операций
func main() {

	push(5)
	push(3)
	result := pop() + pop()
	push(result)

	fmt.Println(stack[0]) // 8
}
