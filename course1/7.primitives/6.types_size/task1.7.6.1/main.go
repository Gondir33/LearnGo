package main

import (
	"unsafe"
)

func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}
func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}
