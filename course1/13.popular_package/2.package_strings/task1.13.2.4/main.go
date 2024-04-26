package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789")
	var sb [4]strings.Builder
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sb[i].WriteRune(chars[rand.Intn(len(chars))])
		}
	}
	ans := make([]string, 0, 10)
	for i := 0; i < 4; i++ {
		ans = append(ans[:], sb[i].String())
	}
	return strings.Join(ans, "-")
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
}
