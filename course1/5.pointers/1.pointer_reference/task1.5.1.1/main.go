package main

func Add(a, b int) *int {
	var sum *int = new(int)
	*sum = a + b
	return sum
}

func Max(numbers []int) *int {
	var max *int = new(int)

	*max = -2147483648
	for i := 0; i < len(numbers); i++ {
		if *max < numbers[i] {
			max = &(numbers[i])
		}
	}
	return max
}
func IsPrime(number int) *bool {
	var ans *bool = new(bool)
	*ans = true
	if number <= 1 {
		*ans = false
		return ans
	}
	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			*ans = false
			break
		}
	}
	return ans
}
func ConcatenateStrings(strs []string) *string {
	var s *string = new(string)

	for i := 0; i < len(strs); i++ {
		*s += strs[i]
	}
	return s
}
