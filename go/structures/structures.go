package structures

import (
	"fmt"
)

// ListNode is implementation of linked list node
type ListNode struct {
	Data     int //For example
	NextNode *ListNode
}

// LinkedList is implementation of linked list data structure
type LinkedList struct {
	Name     string
	HeadNode *ListNode
	LastNode *ListNode
}

// CreateLinkedList is implementation of linked list data structure
func CreateLinkedList(name string) *LinkedList {
	return &LinkedList{
		Name: name,
	}
}

//PushItemToList is implementation of adding new item to linked list
func (list *LinkedList) PushItemToList(data int) {
	item := &ListNode{
		Data: data,
	}
	if list.HeadNode == nil {
		list.HeadNode = item
		list.LastNode = list.HeadNode
	} else {
		list.LastNode = item
		if list.HeadNode.NextNode == nil {
			list.HeadNode.NextNode = list.LastNode
		} else {
			currentNode := list.HeadNode
			for currentNode.NextNode != nil {
				currentNode = currentNode.NextNode
			}
			currentNode.NextNode = list.LastNode
		}
	}
}

// GetLinkedList is implementation of getting all linked list
func GetLinkedList(list *LinkedList) {
	currentNode := list.HeadNode
	if currentNode == nil {
		fmt.Print("\n Linked list is empty")
	} else {
		fmt.Print("\n Head node of linked list: ", currentNode, "\n")
	}
	for currentNode.NextNode != nil {
		currentNode = currentNode.NextNode
		if currentNode.NextNode != nil {
			fmt.Print("\n Linked list current node: ", currentNode, "\n")
		} else {
			fmt.Print("\n Last node of linked list: ", currentNode, "\n")
		}
	}
}

// GetElementOfList is implementation of getting element of linked list by his data
func (list *LinkedList) GetElementOfList(data int) {
	currentNode := list.HeadNode
	if currentNode == nil {
		fmt.Print("Linked list is empty")
	} else {
		if currentNode.Data == data {
			fmt.Print("\n Element ", data, " was founded in ", currentNode, " node \n")
		}
		for currentNode.Data != data {
			currentNode = currentNode.NextNode
			if currentNode == nil {
				fmt.Print("\n there is no ", data, " in this linked list \n")
				break
			}
			if currentNode.Data == data {
				fmt.Print("\n Element ", data, " was founded in ", currentNode, " node \n")
			}
		}
	}
}
