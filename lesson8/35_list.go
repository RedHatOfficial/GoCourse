package main

import "fmt"

// Datový typ představující lineární jednosměrně vázaný seznam
type List[T any] struct {
	Head *Item[T]
}

// Prvek jednosměrně vázaného seznamu
type Item[T any] struct {
	Value T
	Next  *Item[T]
}

// Konstrukce prázdného seznamu (s ukazatelem na nil)
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// Přidání nového prvku na začátek seznamu
func (list *List[T]) Insert(value T) {
	item := Item[T]{
		Value: value,
	}

	// navázání na původní hlavu seznamu
	item.Next = list.Head

	// změna pozice hlavy seznamu
	list.Head = &item
}

// Tisk obsahu celého seznamu
func (list *List[Value]) Print() {
	// první prvek v seznamu (nebo nil)
	item := list.Head

	// postupný průchod dalšími navázanými prvky
	for item != nil {
		fmt.Println(item.Value)

		// přechod na další navázaný prvek
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
}
