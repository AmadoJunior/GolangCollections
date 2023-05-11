package linked_list

type Node[T comparable] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T comparable] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{Length: 0}
}

func (l *LinkedList[T]) Prepend(val T) {
	newNode := &Node[T]{Val: val}
	l.Length += 1
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.next = l.Head
		l.Head.prev = newNode
		l.Head = newNode
	}
}

func (l *LinkedList[T]) Append(val T) {
	newNode := &Node[T]{Val: val}
	l.Length += 1
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		newNode.prev = l.Tail
		l.Tail = newNode
	}
}

func (l *LinkedList[T]) Pop() {
	if l.Tail == nil {
		return
	}
	l.Length -= 1
	l.Tail = l.Tail.prev
	l.Tail.next = nil
}

func (l *LinkedList[T]) Shift() {
	if l.Head == nil {
		return
	}
	l.Length -= 1
	l.Head = l.Head.next
	l.Tail.prev = nil
}

func (l *LinkedList[T]) GetValue(index int) *Node[T] {
	currValue := l.Head
	for i := 0; currValue != nil && i < index; i++ {
		currValue = currValue.next
	}
	return currValue
}

func (l *LinkedList[T]) FindValue(value T) *Node[T] {
	currValue := l.Head
	for currValue != nil && currValue.Val != value {
		currValue = currValue.next
	}
	return currValue
}
