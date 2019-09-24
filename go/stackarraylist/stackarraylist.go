package main

import (
	alist "algorithms/go/arraylist"
	"fmt"
)

// PushToArrayListStack was implemented in arraylist.go as AddNewItemToArrayList

// PopFromArrayListStack is implementation of removing last item from array list based stack
func PopFromArrayListStack(arrayList *alist.ArrayList) {
	fmt.Print("\n array list before pop: ", arrayList)
	arrayList.BaseArray = append(arrayList.BaseArray[:len(arrayList.BaseArray)-1])
	arrayList.Length--
	fmt.Print("\n array list after pop: ", arrayList)
}

func main() {
	x := alist.CreateArrayList()
	fmt.Print("\n new x: ", x)
	x.AddItemToArrayList(1)
	x.AddItemToArrayList(2)
	x.AddItemToArrayList(3)
	x.AddItemToArrayList(4)
	x.AddItemToArrayList(5)
	x.AddItemToArrayList(6)
	x.AddItemToArrayList(7)
	x.AddItemToArrayList(8)
	x.AddItemToArrayList(9)
	x.AddItemToArrayList(10)
	x.AddItemToArrayList(11)
	fmt.Print("\n x with new based array: ", x)
	PopFromArrayListStack(x)
}

/* Simple array implementation

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
*/
