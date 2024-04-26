package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestALL(t *testing.T) {
	want := []int{11, 12, 22, 25, 34, 64, 90}
	data := []int{64, 34, 25, 12, 22, 11, 90}

	sortedData := mergeSort(data)
	assert.Equal(t, want, sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	assert.Equal(t, want, data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	assert.Equal(t, want, data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quicksort(data)
	assert.Equal(t, want, sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	GeneralSort(data)
	assert.Equal(t, want, data)
}
