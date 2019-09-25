package main

import "fmt"

func sum(x []int) int {
	if len(x) == 0 {
		return 0
	}
	return x[0] + sum(append(x[1:]))
}

func countNumberOfItems(x []int) int {
	if len(x) == 0 {
		return 0
	}
	return 1 + countNumberOfItems(append(x[1:]))
}

func maximumNumberOfList(x []int) int {
	if len(x) == 2 {
		if x[0] > x[1] {
			return x[0]
		}
		return x[1]
	}
	subMax := maximumNumberOfList(append(x[1:]))
	if x[0] > subMax {
		return x[0]
	}
	return subMax
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("\n all summ: ", sum(x))
	linearSum := 0
	for i := 0; i < len(x); i++ {
		linearSum += x[i]
	}
	fmt.Print("\n linearSum: ", linearSum)
	numberOfelements := countNumberOfItems(x)
	fmt.Print("\n numberOfelements: ", numberOfelements)
	linearNumberOfElements := 0
	for i := 0; i <= len(x); i++ {
		linearNumberOfElements = i
	}
	fmt.Print("\n linearNumberOfElements: ", linearNumberOfElements)
	maxNumber := maximumNumberOfList(x)
	fmt.Print("\n maxNumber: ", maxNumber)
	linearMaxNumber := 0
	for i := 0; i < len(x); i++ {
		if linearMaxNumber < x[i] {
			linearMaxNumber = x[i]
		}
	}
	fmt.Print("\n linearMaxNumber: ", linearMaxNumber)
}
