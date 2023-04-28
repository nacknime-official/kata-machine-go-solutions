package singlylinkedlist

import dsa "github.com/nacknime-official/kata-machine-go/src/DSA"

type node[T comparable] struct {
	value T
	next  *node[T]
}

type SinglyLinkedList[T comparable] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Prepend(item T) {
	l.length++

	node := &node[T]{value: item}
	if l.length == 1 {
		l.head, l.tail = node, node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *SinglyLinkedList[T]) InsertAt(item T, idx int) {
	if idx < 0 || idx > l.length {
		return
	}

	if idx == l.length {
		l.Append(item)
		return
	} else if idx == 0 {
		l.Prepend(item)
		return
	}
	node := &node[T]{value: item}

	prevNode := l.traverse(idx - 1)
	node.next = prevNode.next
	prevNode.next = node

	l.length++
}

func (l *SinglyLinkedList[T]) Append(item T) {
	l.length++

	node := &node[T]{value: item}
	if l.length == 1 {
		l.head, l.tail = node, node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	if l.length == 0 {
		return dsa.ZeroValue[T](), false
	}

	currNode := l.head
	for i := 0; i < l.length; i++ {
		if currNode.value == item {
			return l.RemoveAt(i)
		}
		currNode = currNode.next
	}
	return dsa.ZeroValue[T](), false
}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	if idx < 0 || idx >= l.length {
		return dsa.ZeroValue[T](), false
	}
	return l.traverse(idx).value, true
}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	if idx < 0 || idx >= l.length {
		return dsa.ZeroValue[T](), false
	}

	if l.length == 1 {
		removingNode := l.head
		l.head, l.tail = nil, nil
		l.length--
		return removingNode.value, true
	}

	prevNode := l.traverse(idx - 1)
	removingNode := prevNode.next

	// remove first
	if idx == 0 {
		removingNode = prevNode
		l.head = removingNode.next
	} else if idx == l.length-1 { // remove last
		prevNode.next = nil
		l.tail = prevNode
	} else { // remove somewhere in the middle
		prevNode.next = removingNode.next
	}
	removingNode.next = nil

	l.length--
	return removingNode.value, true
}

func (l *SinglyLinkedList[T]) traverse(idx int) *node[T] {
	if idx == l.length-1 {
		return l.tail
	}

	currNode := l.head
	for i := 0; i < idx; i++ {
		currNode = currNode.next
	}
	return currNode
}
