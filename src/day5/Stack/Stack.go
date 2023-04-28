package stack

import dsa "github.com/nacknime-official/kata-machine-go/src/DSA"

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	Length int
	head   *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.Length++
	node := &node[T]{value: item}

	node.next = s.head
	s.head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length == 0 {
		return dsa.ZeroValue[T](), false
	}
	poppingNode := s.head

	s.head = poppingNode.next
	poppingNode.next = nil

	s.Length--
	return poppingNode.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Length == 0 {
		return dsa.ZeroValue[T](), false
	}
	return s.head.value, true
}
