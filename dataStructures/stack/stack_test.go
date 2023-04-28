package stack
import("testing")

func TestStack(t *testing.T) {
	myStack := New()
	if myStack.Len() != 0 {
		t.Errorf("Expected length of new stack to be 0, got %d", myStack.Len())
	}

	myStack.Push(1)
	if myStack.Len() != 1 {
		t.Errorf("Expected length of stack after push to be 1, got %d", myStack.Len())
	}

	if myStack.Peek() != 1 {
		t.Errorf("Expected top value of stack to be 1, got %v", myStack.Peek())
	}

	val := myStack.Pop()
	if val != 1 {
		t.Errorf("Expected popped value to be 1, got %v", val)
	}

	if myStack.Len() != 0 {
		t.Errorf("Expected length of stack after pop to be 0, got %d", myStack.Len())
	}
}
