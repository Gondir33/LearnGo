package main

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	bytes := []byte(*str)

	var tmp byte
	for i := 0; i < len(bytes)/2; i++ {
		tmp = bytes[i]
		bytes[i] = bytes[len(bytes)-i-1]
		bytes[len(bytes)-i-1] = tmp
	}
	*str = string(bytes[:])
}
