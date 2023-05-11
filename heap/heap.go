package heap

type Heap struct {
	data       []int
	comparator func(int, int) bool
}

func NewHeapFromSlice(slice []int, comparator func(int, int) bool) *Heap {
	newHeap := &Heap{data: slice, comparator: comparator}
	newHeap.Heapify()
	return newHeap
}

func NewHeap(comparator func(int, int) bool) *Heap {
	return &Heap{data: make([]int, 0), comparator: comparator}
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

func (h *Heap) bubbleUp() {
	currIndex := len(h.data) - 1
	parentIndex := getParentIndex(currIndex)

	for parentIndex >= 0 && h.comparator(h.data[currIndex], h.data[parentIndex]) {
		h.data[currIndex], h.data[parentIndex] = h.data[parentIndex], h.data[currIndex]
		currIndex = parentIndex
		parentIndex = getParentIndex(currIndex)
	}
}

func (h *Heap) bubbleDown(startIndex int) {
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

func (h *Heap) Push(num int) {
	h.data = append(h.data, num)
	h.bubbleUp()
}

func (h *Heap) Pop() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(0)
}

func (h *Heap) GetTop() int {
	return h.data[0]
}

func (h *Heap) Heapify() {
	lastNonLeafIndex := getParentIndex(len(h.data) - 1)

	for i := lastNonLeafIndex; i >= 0; i-- {
		h.bubbleDown(i)
	}
}
