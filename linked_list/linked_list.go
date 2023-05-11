package linked_list

type Node struct {
	val  int
	next *Node
	prev *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{length: 0}
}

func (l *LinkedList) Prepend(val int) {
	newNode := &Node{val: val}
	l.length += 1
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{val: val}
	l.length += 1
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

func (l *LinkedList) Pop() {
	if l.tail == nil {
		return
	}
	l.length -= 1
	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *LinkedList) Shift() {
	if l.head == nil {
		return
	}
	l.length -= 1
	l.head = l.head.next
	l.tail.prev = nil
}

func (l *LinkedList) GetLen() int {
	return l.length
}

func (l *LinkedList) GetValue(index int) *Node {
	currValue := l.head
	for i := 0; currValue != nil && i < index; i++ {
		currValue = currValue.next
	}
	return currValue
}

func (l *LinkedList) FindValue(value int) *Node {
	currValue := l.head
	for currValue != nil && currValue.val != value {
		currValue = currValue.next
	}
	return currValue
}
