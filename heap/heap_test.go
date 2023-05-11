package heap

import (
	"testing"
)

func TestPushMinHeap(t *testing.T) {
	minHeap := NewHeap(func(x int, y int) bool { return x < y })
	minHeap.Push(5)
	minHeap.Push(1)
	minHeap.Push(0)
	minHeap.Push(10)
	minValue := minHeap.GetTop()
	if minValue != 0 {
		t.Errorf("minHeap.GetTop() FAILED. Expected %d, Got %d\n", 0, minValue)
	} else {
		t.Logf("minHeap.GetTop() PASSED. Expected %d, Got %d\n", 0, minValue)
	}
}

func TestPopMinHeap(t *testing.T) {
	minHeap := NewHeap(func(x int, y int) bool { return x < y })
	minHeap.Push(5)
	minHeap.Push(1)
	minHeap.Push(0)
	minHeap.Push(10)
	minHeap.Pop()
	minValue := minHeap.GetTop()
	if minValue != 1 {
		t.Errorf("minHeap.GetTop() FAILED. Expected %d, Got %d\n", 1, minValue)
	} else {
		t.Logf("minHeap.GetTop() PASSED. Expected %d, Got %d\n", 1, minValue)
	}
}

func TestPushMaxHeap(t *testing.T) {
	maxHeap := NewHeap(func(x int, y int) bool { return x > y })
	maxHeap.Push(5)
	maxHeap.Push(1)
	maxHeap.Push(0)
	maxHeap.Push(10)
	minValue := maxHeap.GetTop()
	if minValue != 10 {
		t.Errorf("maxHeap.GetTop() FAILED. Expected %d, Got %d\n", 10, minValue)
	} else {
		t.Logf("maxHeap.GetTop() PASSED. Expected %d, Got %d\n", 10, minValue)
	}
}

func TestPopMaxHeap(t *testing.T) {
	maxHeap := NewHeap(func(x int, y int) bool { return x > y })
	maxHeap.Push(5)
	maxHeap.Push(1)
	maxHeap.Push(0)
	maxHeap.Push(10)
	maxHeap.Pop()
	minValue := maxHeap.GetTop()
	if minValue != 5 {
		t.Errorf("maxHeap.GetTop() FAILED. Expected %d, Got %d\n", 5, minValue)
	} else {
		t.Logf("maxHeap.GetTop() PASSED. Expected %d, Got %d\n", 5, minValue)
	}
}

func TestNewMinHeapFromSlice(t *testing.T) {
	sampleSlice := []int{5, 1, 0, 10}
	minHeap := NewHeapFromSlice(sampleSlice, func(x int, y int) bool { return x < y })
	minValue := minHeap.GetTop()
	if minValue != 0 {
		t.Errorf("minHeap.GetTop() FAILED. Expected %d, Got %d\n", 0, minValue)
	} else {
		t.Logf("minHeap.GetTop() PASSED. Expected %d, Got %d\n", 0, minValue)
	}
}

func TestNewMaxHeapFromSlice(t *testing.T) {
	sampleSlice := []int{5, 1, 0, 10}
	maxHeap := NewHeapFromSlice(sampleSlice, func(x int, y int) bool { return x > y })
	minValue := maxHeap.GetTop()
	if minValue != 10 {
		t.Errorf("maxHeap.GetTop() FAILED. Expected %d, Got %d\n", 10, minValue)
	} else {
		t.Logf("maxHeap.GetTop() PASSED. Expected %d, Got %d\n", 10, minValue)
	}
}
