package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[(len(arr)-1)/2]
	fmt.Print("\n Selected pivot: ", pivot, "\n")
	var less []int
	var greater []int
	for i := range arr {
		if arr[i] < pivot {
			less = append(less, arr[i])
		}
		if arr[i] == pivot {
			continue
		}
		if arr[i] > pivot {
			greater = append(greater, arr[i])
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, pivot), greater...)
}

func main() {
	testSlice := []int{2, 5, 10, 3, 4, 28, 45, 37, 1}
	fmt.Print("\n test slice before sorting: ", testSlice, "\n")
	fmt.Print("\n test slice after sorting: ", quickSort(testSlice), "\n")
}
