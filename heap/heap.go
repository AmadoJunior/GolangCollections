package heap

type Heap[T comparable] struct {
	data       []T
	comparator func(T, T) bool
}

func NewHeapFromSlice[T comparable](slice []T, comparator func(T, T) bool) *Heap[T] {
	newHeap := &Heap[T]{data: slice, comparator: comparator}
	newHeap.Heapify()
	return newHeap
}

func NewHeap[T comparable](comparator func(T, T) bool) *Heap[T] {
	return &Heap[T]{data: make([]T, 0), comparator: comparator}
}

func getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}

func getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}

func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap[T]) bubbleUp() {
	currIndex := len(h.data) - 1
	parentIndex := getParentIndex(currIndex)

	for parentIndex >= 0 && h.comparator(h.data[currIndex], h.data[parentIndex]) {
		h.data[currIndex], h.data[parentIndex] = h.data[parentIndex], h.data[currIndex]
		currIndex = parentIndex
		parentIndex = getParentIndex(currIndex)
	}
}

func (h *Heap[T]) bubbleDown(startIndex int) {
	currIndex := startIndex
	leftChildIndex := getLeftChildIndex(currIndex)
	rightChildIndex := getRightChildIndex(currIndex)
	maxIndex := currIndex

	if leftChildIndex < len(h.data) && h.comparator(h.data[leftChildIndex], h.data[maxIndex]) {
		maxIndex = leftChildIndex
	}

	if rightChildIndex < len(h.data) && h.comparator(h.data[rightChildIndex], h.data[maxIndex]) {
		maxIndex = rightChildIndex
	}

	if maxIndex != currIndex {
		h.data[currIndex], h.data[maxIndex] = h.data[maxIndex], h.data[currIndex]
		h.bubbleDown(maxIndex)
	}
}

func (h *Heap[T]) Push(num T) {
	h.data = append(h.data, num)
	h.bubbleUp()
}

func (h *Heap[T]) Pop() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(0)
}

func (h *Heap[T]) GetTop() T {
	return h.data[0]
}

func (h *Heap[T]) Heapify() {
	lastNonLeafIndex := getParentIndex(len(h.data) - 1)

	for i := lastNonLeafIndex; i >= 0; i-- {
		h.bubbleDown(i)
	}
}
