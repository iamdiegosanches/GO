package main

import (
	"fmt"
)

type Node struct {
	next *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (L *List) GetSize() int {
	return L.size
}

func (L *List) Insert(value interface{}) {
	list := &Node{
		next: nil,
		value: value,
	}
	if L.head == nil {
		L.head = list
		L.tail = list
	} else {
		L.tail.next = list
		L.tail = list
	}
	L.size++
}

func (L *List) Display() {
	list := L.head
	for i := 0; i < L.GetSize(); i++ {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
}

func (L *List) ReverseList() {
	current := L.head
	previous, next := &Node{}, &Node{}
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	L.head = previous
}

// TODO: implement search function
func (L *List) Search(val interface{})  {
	panic("Unimplemented function")
}

// TODO: implement delete function
func (L *List) Delete(val interface{}) {
	panic("Unimplemented function")
}

func main() {
	myList := List{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Insert(5)

	myList.ReverseList()

	myList.Display()

}
