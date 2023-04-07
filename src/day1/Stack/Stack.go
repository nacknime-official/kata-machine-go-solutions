package stack

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
	s.head = &node[T]{value: item, next: s.head}
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length == 0 {
		return zeroValue[T](), false
	}

	poppingNode := s.head
	s.head = poppingNode.next
	poppingNode.next = nil

	s.Length--
	return poppingNode.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Length == 0 {
		return zeroValue[T](), false
	}
	return s.head.value, true
}

func zeroValue[T any]() T {
	var zero T
	return zero
}
