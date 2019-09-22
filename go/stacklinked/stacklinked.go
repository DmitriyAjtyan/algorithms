package main

import (
	str "algorithms/go/structures"
	"fmt"
)

// PushToStack is implementation of adding new item to stack
func PushToStack(stack str.LinkedList, newData int) {
	fmt.Print("\n stack last node before push: ", stack.LastNode)
	stack.PushItemToList(newData)
	fmt.Print("\n stack last node after push: ", stack.LastNode)
}

// PopFromStack is implementation of removing last added item from stack
func PopFromStack(stack *str.LinkedList) {
	fmt.Print("\n stack last node before pop: ", stack.LastNode)
	currentNode := stack.HeadNode
	for currentNode.NextNode != stack.LastNode {
		currentNode = currentNode.NextNode
	}
	fmt.Print("\n stack pre last node is: ", currentNode)
	currentNode.NextNode = nil
	stack.LastNode = currentNode
	fmt.Print("\n stack last node after pop: ", stack.LastNode)
}

func main() {
	// linked list stack test
	newStack := str.CreateLinkedList("newStack")
	newStack.PushItemToList(1)
	newStack.PushItemToList(2)
	newStack.PushItemToList(3)
	newStack.PushItemToList(4)
	newStack.PushItemToList(5)
	str.GetLinkedList(newStack)
	PopFromStack(newStack)
	str.GetLinkedList(newStack)
}
