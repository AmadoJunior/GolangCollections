package linked_list

import "testing"

func TestPrepend(t *testing.T) {
	list := NewLinkedList[int]()

	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)

	// Check the length of the list
	expectedLen := 3
	if list.Length != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.Length)
	}

	// Check the values in the list
	expectedValues := []int{1, 2, 3}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.Val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.Val)
		}
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Check the length of the list
	expectedLen := 3
	if list.Length != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.Length)
	}

	// Check the values in the list
	expectedValues := []int{1, 2, 3}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.Val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.Val)
		}
	}
}

func TestFindValue(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Find a value in the list
	value := 2
	expectedValue := 2
	node := list.FindValue(value)
	if node == nil {
		t.Errorf("Expected to find value: %d, but got nil", value)
	} else if node.Val != expectedValue {
		t.Errorf("Expected value: %d, but got: %d", expectedValue, node.Val)
	}

	// Try to find a non-existing value
	nonExistingValue := 5
	node = list.FindValue(nonExistingValue)
	if node != nil {
		t.Errorf("Expected nil, but got a node with value: %d", node.Val)
	}
}

func TestPop(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Pop()

	// Check the length of the list
	expectedLen := 2
	if list.Length != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.Length)
	}

	// Check the values in the list
	expectedValues := []int{1, 2}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.Val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.Val)
		}
	}
}

func TestShift(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Shift()

	// Check the length of the list
	expectedLen := 2
	if list.Length != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.Length)
	}

	// Check the values in the list
	expectedValues := []int{2, 3}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.Val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.Val)
		}
	}
}
