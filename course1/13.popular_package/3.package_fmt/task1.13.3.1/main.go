package main

import "fmt"

func generateMathString(operands []int, operator string) string {
	ans := operands[0]
	switch operator {
	case "+":
		for i := 1; i < len(operands); i++ {
			ans += operands[i]
		}
	case "-":
		for i := 1; i < len(operands); i++ {
			ans -= operands[i]
		}
	case "/":
		for i := 1; i < len(operands); i++ {
			ans -= operands[i]
		}
	case "*":
		for i := 1; i < len(operands); i++ {
			ans -= operands[i]
		}
	}
	var s string
	for i := 0; i < len(operands)-1; i++ {
		s += fmt.Sprintf("%v %v ", operands[i], operator)
	}
	return fmt.Sprintf("%v%v = %v", s, operands[len(operands)-1], ans)

}

// Пример результата выполнения программы:
func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "-")) // "2 + 4 + 6 = 12"
}
