package linkedlist

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

func (L *List) GetSize() int {
	return L.size
}

func (L *List) Insert(val interface{}) {
	list := &Node{
		next: nil,
		value: val,
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

func (L *List) Search(val interface{}) interface{} {
	list := L.head
	for i := 0; i < L.GetSize(); i++ {
		if list.value == val {
			return list.value
		}
		list = list.next
	}
	return nil
}

func (L *List) Delete(val interface{}) bool {
	prev := L.head
	curr := L.head.next

	for curr != nil {
		if curr.value == val {
			prev.next = curr.next
			if L.tail == curr {
				L.tail = prev
			}
			L.size--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}
