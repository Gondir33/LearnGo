package main

func Factorial(n *int) int {
	res := 1
	for i := *n; i > 1; i-- {
		res *= i
	}
	return res
}

func isPalindrome(str *string) bool {
	bytes := []byte(*str)
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-i-1] {
			return false
		}
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	trg := *target
	nums := *numbers
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == trg {
			ans++
		}
	}
	return ans
}

func ReverseString(str *string) string {
	bytes := []byte(*str)
	var tmp byte
	for i := 0; i < len(bytes)/2; i++ {
		tmp = bytes[i]
		bytes[i] = bytes[len(bytes)-i-1]
		bytes[len(bytes)-i-1] = tmp
	}
	return string(bytes[:])
}
