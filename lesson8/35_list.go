package main

import "fmt"

type User struct {
	id      int
	name    string
	surname string
}

type allowed_type interface {
	~int | ~string | User
}

// Singly linked list header structure
type List[T allowed_type] struct {
	Head *Item[T]
}

// Item stored in singly linked list
type Item[T allowed_type] struct {
	Value T
	Next  *Item[T]
}

// NewList constructs empty singly linked list
func NewList[T allowed_type]() *List[T] {
	return &List[T]{}
}

// Insert method adds new item into the list
func (list *List[T]) Insert(value T) {
	item := Item[T]{
		Value: value,
	}

	// link to the recend head
	item.Next = list.Head

	// change head to point/link to new item
	list.Head = &item
}

// Print method prints the whole list
func (list *List[Value]) Print() {
	// reference to the first list item (or nil)
	item := list.Head

	// iterate over all list items
	for item != nil {
		fmt.Println(item.Value)

		// go to the next item in the list
		item = item.Next
	}

}

func main() {
	list1 := NewList[int]()

	list1.Insert(1)
	list1.Insert(2)
	list1.Insert(3)
	list1.Insert(4)

	list1.Print()

	fmt.Println()
	list2 := NewList[string]()

	list2.Insert("first")
	list2.Insert("second")
	list2.Insert("third")
	list2.Insert("fourth")

	list2.Print()

	user1 := User{
		id:      1,
		name:    "John",
		surname: "Doe",
	}

	list3 := NewList[User]()
	list3.Insert(user1)
	list3.Print()

}
