package main

import "fmt"

var i = 5

func countdown(i int) {
	fmt.Print("\ni: ", i)
	if i != 0 {
		countdown(i - 1)
	}
}

func factorial(i int) int {
	if i == 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	countdown(i)
	i = factorial(i)
	fmt.Print("\n i factorial: ", i)
}
