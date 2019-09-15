package linked

import "fmt"

type listItem struct {
	data     int //For example
	nextItem *listItem
}

// array is implementation of array data structure
type array struct {
	firstItem  int
	secondItem int
	thirdItem  int
	fourthItem int
}

// List is implementation of linked list data structure
func List() {
	var headItem listItem
	var temporary listItem
	var newItem listItem
	headItem.data = 1
	temporary = headItem
	fmt.Print("\ntemporary first position: ", temporary, "\n")
	newItem.data = 2
	newItem.nextItem = nil
	headItem.nextItem = &newItem
	temporary = newItem
	fmt.Print("\ntemporary second position: ", temporary, "\n")
	fmt.Print("\nhead item: ", headItem, "\n")
	fmt.Print("\nnew/last item: ", newItem, "\n")

	var newArray array
	newArray.firstItem = 0
	newArray.secondItem = 1
	newArray.thirdItem = 2
	newArray.fourthItem = 3
}
