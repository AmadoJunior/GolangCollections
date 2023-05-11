package binary_tree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the in-order traversal result
	var result []int
	root.InOrderTraversal(&result)

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert: Expected in-order traversal result: %v, but got: %v", expected, result)
	}
}

func TestSearch(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Search for existing and non-existing values
	existingValue := 4
	nonExistingValue := 10

	existingNode := root.Search(existingValue)
	if existingNode == nil || existingNode.val != existingValue {
		t.Errorf("Search: Failed to find existing value: %d", existingValue)
	}

	nonExistingNode := root.Search(nonExistingValue)
	if nonExistingNode != nil {
		t.Errorf("Search: Found non-existing value: %d", nonExistingValue)
	}
}

func TestMin(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	min := root.Min()
	expectedMin := 2
	if min != expectedMin {
		t.Errorf("Min: Expected minimum value: %d, but got: %d", expectedMin, min)
	}
}

func TestMax(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	max := root.Max()
	expectedMax := 8
	if max != expectedMax {
		t.Errorf("Max: Expected maximum value: %d, but got: %d", expectedMax, max)
	}
}

func TestInOrderTraversal(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the in-order traversal result
	var result []int
	root.InOrderTraversal(&result)

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrderTraversal: Expected in-order traversal result: %v, but got: %v", expected, result)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	root := MakeBST(5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the pre-order traversal result
	var result []int
	root.PreOrderTraversal(&result)

	expected := []int{5, 3, 2, 4, 7, 6, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderTraversal: Expected pre-order travelsal result: %v, but got: %v", expected, result)
	}
}
