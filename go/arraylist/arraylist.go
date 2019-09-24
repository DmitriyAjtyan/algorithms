package arraylist

// ArrayList is implementation of array list data structure
type ArrayList struct {
	Capacity  uint
	Length    uint
	BaseArray []int
}

// CreateArrayList is implementation of creating array list data structure
func CreateArrayList() *ArrayList {
	return &ArrayList{
		Capacity: 10,
		Length:   0,
	}
}

// AddItemToArrayList is implementation of adding item to array list
func (arrayList *ArrayList) AddItemToArrayList(newItem int) {
	if arrayList.Capacity == arrayList.Length {
		arrayList.Capacity *= 2
		newBaseArray := make([]int, 0, arrayList.Capacity)
		newBaseArray = append(newBaseArray[:0], arrayList.BaseArray...)
		arrayList.BaseArray = newBaseArray
	}
	newSlice := []int{newItem}
	arrayList.BaseArray = append(arrayList.BaseArray[:arrayList.Length], newSlice...)
	arrayList.Length++
}

// GetElementOfArrayList is implementation of getting element of array list data structure by index
func GetElementOfArrayList(arrayList *ArrayList, index uint) int {
	return arrayList.BaseArray[index]
}

/*
func main() {
	x := CreateArrayList()
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
	element := GetElementOfArrayList(x, 10)
	fmt.Print("\n element of list: ", element)
}
*/
