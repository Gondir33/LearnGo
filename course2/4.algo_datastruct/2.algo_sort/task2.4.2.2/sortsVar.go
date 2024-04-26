package main

func GeneralSort(arr []int) {
	quicksort(arr)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func quicksort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

/*
func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)

	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quicksort(data)
	fmt.Println("Sorted by Quicksort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", data)
}
*/
