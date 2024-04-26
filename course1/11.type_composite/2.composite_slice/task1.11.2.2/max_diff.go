package main

func MaxDifference(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}
	var max, min int
	max, min = numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return max - min
}
