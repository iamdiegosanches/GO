package stack

type Node struct {
	next *Node
	value interface{}
}

type Stack struct {
	top *Node
	bottom *Node
	size int
}

func New() *Stack {
	s := new(Stack)
	return s
}

func (this *Stack) Len() int {
	return this.size
}

func (this *Stack) Peek() interface{} {
	return this.top.value
}

func (this *Stack) Pop() interface{} {
	if this.top == nil {
		return nil
	}
	val := this.top.value
	this.top = this.top.next
	this.size--
	if this.size == 0 {
		this.bottom = nil
	}
	return val
}

func (this *Stack) Push(value interface{}) {
	newNode := &Node{
		value: value,
	}
	if this.bottom == nil {
		this.bottom = newNode
		this.top = newNode
	} else {
		this.top.next = newNode
		this.top = newNode
	}
	this.size++
}
