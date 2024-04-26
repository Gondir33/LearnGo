package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	return "Name: " + fake.FullName() +
		"\nAddress: " + fake.StreetAddress() +
		"\nPhone: " + fake.Phone() +
		"\nEmail: " + fake.EmailAddress()
}

func main() {
	fmt.Println(GenerateFakeData())
}
