package main

import "fmt"

var (
	list [8]int
	item int
)

func binarySearch(list [8]int, item int) int {
	low := 0
	high := len(list) - 1
	mid := (high - low) / 2
	guess := list[mid]

	for low <= high {
		mid = (high + low) / 2
		guess = list[mid]
		if guess < item {
			low = mid + 1
		} else if guess > item {
			high = mid - 1
		} else {
			break
		}
	}
	return guess
}

func main() {
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	item = 0
	fmt.Print("\nbinary search result: ", binarySearch(list, item), "\n")
}
