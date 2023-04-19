package queue

import dsa "github.com/nacknime-official/kata-machine-go/src/DSA"

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	Length int
	head   *node[T]
	tail   *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.Length++
	node := &node[T]{value: item}

	if q.Length == 1 {
		q.head, q.tail = node, node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.Length == 0 {
		return dsa.ZeroValue[T](), false
	}
	dequedNode := q.head

	q.head = dequedNode.next
	dequedNode.next = nil

	q.Length--
	return dequedNode.value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		return dsa.ZeroValue[T](), false
	}
	return q.head.value, true
}
