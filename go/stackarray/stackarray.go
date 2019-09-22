package main

import "fmt"

// PushToArrayStack is implementation of adding new item to array stack
func PushToArrayStack(stack []int, newItem int) {
	fmt.Print("\n stack before pushing new item: ", stack)
	newItemForStack := []int{newItem}
	stack = append(stack, newItemForStack...)
	fmt.Print("\n stack after pushing new item: ", stack)
}

// PopFromArrayStack is implementation of removing last item from array stack
func PopFromArrayStack(stack []int) {
	fmt.Print("\n stack before pop: ", stack)
	stack = append(stack[:(len(stack) - 1)])
	fmt.Print("\n stack after pop: ", stack)
}

func main() {
	// array stack test
	newArrayStack := []int{1, 2, 3, 4, 5}
	PushToArrayStack(newArrayStack, 6)
	PopFromArrayStack(newArrayStack)
}
