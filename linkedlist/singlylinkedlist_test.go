package linkedlist
import("testing")

func TestLinkedList(t *testing.T) {
	list := &List{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	if list.GetSize() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.GetSize())
	}

	if list.Search(2) != 2 {
		t.Errorf("Expected to find value 2 in the list")
	}

	if list.Search(4) != nil {
		t.Errorf("Expected to not find value 4 in the list")
	}

	list.ReverseList()

	if list.head.value != 3 {
		t.Errorf("Expected head value to be 3 after reversing the list")
	}

	if !list.Delete(1) {
		t.Errorf("Expected Deletion of element")
	}

	if list.Delete(1) {
		t.Errorf("Expected element not in the list")
	}

	if list.GetSize() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.GetSize())
	}
	
	list.Display()
}
