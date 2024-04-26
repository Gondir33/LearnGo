package main

import "fmt"

func helper(flag int) string {
	if flag == 0 {
		return "-,-,-"
	} else if flag == 1 {
		return "-,-,Execute"
	} else if flag == 2 {
		return "-,Write,-"
	} else if flag == 3 {
		return "_,Write,Execute"
	} else if flag == 4 {
		return "Read,-,-"
	} else if flag == 5 {
		return "Read,-,Execute"
	} else if flag == 6 {
		return "Read,Write,-"
	} else if flag == 7 {
		return "Read,Write,Execute"
	} else {
		return "-,-,-"
	}
}

func getFilePermissions(flag int) string {
	ans := "Owner:"
	ans += helper(flag / 100)
	ans += " Group:"
	ans += helper((flag / 10) % 10)
	ans += " Other:"
	ans += helper(flag % 10)
	return ans
}

func main() {
	fmt.Println(getFilePermissions(700))
	fmt.Println(getFilePermissions(644))
	fmt.Println(getFilePermissions(755))
	fmt.Println(getFilePermissions(0))
}
