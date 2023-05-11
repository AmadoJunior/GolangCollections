package linked_list

import "testing"

func TestPrepend(t *testing.T) {
	list := NewLinkedList()

	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)

	// Check the length of the list
	expectedLen := 3
	if list.GetLen() != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.GetLen())
	}

	// Check the values in the list
	expectedValues := []int{1, 2, 3}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.val)
		}
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Check the length of the list
	expectedLen := 3
	if list.GetLen() != expectedLen {
		t.Errorf("Expected length: %d, but got: %d", expectedLen, list.GetLen())
	}

	// Check the values in the list
	expectedValues := []int{1, 2, 3}
	for i, val := range expectedValues {
		node := list.GetValue(i)
		if node.val != val {
			t.Errorf("Expected value: %d, but got: %d", val, node.val)
		}
	}
}

func TestFindValue(t *testing.T) {
	list := NewLinkedList()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Find a value in the list
	value := 2
	expectedValue := 2
	node := list.FindValue(value)
	if node == nil {
		t.Errorf("Expected to find value: %d, but got nil", value)
	} else if node.val != expectedValue {
		t.Errorf("Expected value: %d, but got: %d", expectedValue, node.val)
	}

	// Try to find a non-existing value
	nonExistingValue := 5
	node = list.FindValue(nonExistingValue)
	if node != nil {
		t.Errorf("Expected nil, but got a node with value: %d", node.val)
	}
}
