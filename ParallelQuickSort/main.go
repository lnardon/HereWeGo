package main

import (
	"fmt"
)

func basicQuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := 0
	pivot := arr[pivotIndex]

	lower := make([]int, 0)
	higher := make([]int, 0)
	for i := range arr {
		if i == pivotIndex {
			continue
		}
		if arr[i] < pivot {
			lower = append(lower, arr[i])
		} else {
			higher = append(higher, arr[i])
		}
	}

	return append(append(basicQuickSort(lower), pivot), basicQuickSort(higher)...)
}

func main() {
	arr := []int{2, 5, 3, 9, 8, 10, 7, 4,6, 1}
	fmt.Println("Unsorted array:", arr)
	sortedArr := basicQuickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
