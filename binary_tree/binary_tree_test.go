package binary_tree

import (
	"reflect"
	"testing"
)

type CustomType int

func (m CustomType) Less(other Comparable) bool {
	otherVal := other.(CustomType)
	return m < otherVal
}

func (m CustomType) More(other Comparable) bool {
	otherVal := other.(CustomType)
	return m > otherVal
}

func (m CustomType) Equal(other Comparable) bool {
	otherVal := other.(CustomType)
	return m == otherVal
}

func TestInsert(t *testing.T) {
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the in-order traversal result
	var result []CustomType
	root.InOrderTraversal(&result)

	expected := []CustomType{2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert: Expected in-order traversal result: %v, but got: %v", expected, result)
	}
}

func TestSearch(t *testing.T) {
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Search for existing and non-existing values
	var existingValue CustomType = 4
	var nonExistingValue CustomType = 10

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
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	var min CustomType = root.Min()
	var expectedMin CustomType = 2
	if !min.Equal(expectedMin) {
		t.Errorf("Min: Expected minimum value: %d, but got: %d", expectedMin, min)
	}
}

func TestMax(t *testing.T) {
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	var max CustomType = root.Max()
	var expectedMax CustomType = 8
	if !max.Equal(expectedMax) {
		t.Errorf("Max: Expected maximum value: %d, but got: %d", expectedMax, max)
	}
}

func TestInOrderTraversal(t *testing.T) {
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the in-order traversal result
	var result []CustomType
	root.InOrderTraversal(&result)

	expected := []CustomType{2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrderTraversal: Expected in-order traversal result: %v, but got: %v", expected, result)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	root := MakeBST[CustomType](5)

	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// Create a slice to store the pre-order traversal result
	var result []CustomType
	root.PreOrderTraversal(&result)

	expected := []CustomType{5, 3, 2, 4, 7, 6, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderTraversal: Expected pre-order travelsal result: %v, but got: %v", expected, result)
	}
}
