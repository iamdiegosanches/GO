package main

import "fmt"

type Node struct {
	next *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
	size int
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
	for list != nil {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
}

func main() {
	myList := List{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)

	myList.Display()
}
