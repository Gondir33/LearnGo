package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := ""
	valid := isValidEmail(email)
	if valid {
		fmt.Printf("%s является валидным email-адресом\n", email)
	} else {
		fmt.Printf("%s не является валидным email-адресом\n", email)
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`@{1}`)
	if re.FindString(email) == "@" {
		return true
	} else {
		return false
	}
}
