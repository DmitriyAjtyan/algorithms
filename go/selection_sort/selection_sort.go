package main

import (
	"fmt"
)

// FindMaxValueFromList is finding largest value from list
func FindMaxValueFromList(list []int) (int, int) {
	max := list[0]
	indexOfMaxElement := 0
	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
			indexOfMaxElement = i
		}
	}
	return max, indexOfMaxElement
}

// FindMinValueFromList is finding smallest value from list
func FindMinValueFromList(list []int) (int, int) {
	min := list[0]
	indexOfMinElement := 0
	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
			indexOfMinElement = i
		}
	}
	return min, indexOfMinElement
}

// SortListBySelectionMethodFromMinToMax is function for list sorting by selection method from min to max value
func SortListBySelectionMethodFromMinToMax(list []int) []int {
	var sortedList []int
	currentMin := []int{0}
	var indexOfCurrentMin int
	for len(list) != 0 {
		currentMin[0], indexOfCurrentMin = FindMinValueFromList(list)
		sortedList = append(sortedList, currentMin...)
		list = append(list[:indexOfCurrentMin], list[indexOfCurrentMin+1:]...)
	}
	return sortedList
}

// SortListBySelectionMethodFromMaxToMin is function for list sorting by selection method from max to min value
func SortListBySelectionMethodFromMaxToMin(list []int) []int {
	var sortedList []int
	currentMin := []int{0}
	var indexOfCurrentMin int
	for len(list) != 0 {
		currentMin[0], indexOfCurrentMin = FindMaxValueFromList(list)
		sortedList = append(sortedList, currentMin...)
		list = append(list[:indexOfCurrentMin], list[indexOfCurrentMin+1:]...)
	}
	return sortedList
}

func main() {
	someList := []int{10, 5, 7, 15, 2, 0, 31}
	someList = SortListBySelectionMethodFromMinToMax(someList)
	fmt.Print("\n\n List after sorting by selection method from min to max: ", someList, "\n")
	someList = SortListBySelectionMethodFromMaxToMin(someList)
	fmt.Print("\n\n List after sorting by selection method form max to min: ", someList, "\n")
}
