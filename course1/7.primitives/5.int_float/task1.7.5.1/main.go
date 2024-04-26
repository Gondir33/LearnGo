package main

import (
	"fmt"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
	var number uint32
	// Преобразование строки в двоичной системе в целочисленное представление
	bytes := []byte(binary)
	fmt.Println(bytes)
	for i := 0; i < len(bytes); i++ {
		number <<= 1
		if bytes[i] == 49 {
			number |= 1
		}
	}
	// Преобразование целочисленного представления в число с плавающей точкой
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}
