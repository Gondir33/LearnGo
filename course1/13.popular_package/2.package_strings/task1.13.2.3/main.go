package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(byte(rand.Intn(128)))
	}
	return sb.String()
}

func main() {
	randomString := GenerateRandomString(10)
	fmt.Println(randomString)
}
