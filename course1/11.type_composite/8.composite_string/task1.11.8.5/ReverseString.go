package main

func ReverseString(str string) string {
	bytes := []byte(str)

	for i := 0; i < len(bytes)/2; i++ {
		tmp := bytes[i]
		bytes[i] = bytes[len(bytes)-1-i]
		bytes[len(bytes)-1-i] = tmp
	}
	return string(bytes)
}
