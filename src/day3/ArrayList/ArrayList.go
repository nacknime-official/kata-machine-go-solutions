package arraylist

import dsa "github.com/nacknime-official/kata-machine-go/src/DSA"

type ArrayList[T comparable] struct {
	length int
	cap    int
	arr    []T
}

func (l *ArrayList[T]) Len() int { return l.length }

func NewArrayList[T comparable](capacity int) *ArrayList[T] {
	return &ArrayList[T]{
		length: 0,
		cap:    capacity,
		arr:    make([]T, capacity),
	}
}

func (l *ArrayList[T]) Prepend(item T) {
	l.InsertAt(item, 0)
}

func (l *ArrayList[T]) InsertAt(item T, idx int) {
	if idx > l.length {
		return
	}

	if l.length == l.cap {
		l.increaseArrSize()
	}

	l.shiftRight(idx)
	l.arr[idx] = item
	l.length++
}

func (l *ArrayList[T]) Append(item T) {
	l.InsertAt(item, l.length)
}

func (l *ArrayList[T]) Remove(item T) (T, bool) {
	for i := 0; i < l.length; i++ {
		if item == l.arr[i] {
			return l.RemoveAt(i)
		}
	}
	return dsa.ZeroValue[T](), false
}

func (l *ArrayList[T]) Get(idx int) (T, bool) {
	if idx >= l.length {
		return dsa.ZeroValue[T](), false
	}
	return l.arr[idx], true
}

func (l *ArrayList[T]) RemoveAt(idx int) (T, bool) {
	if idx >= l.length {
		return dsa.ZeroValue[T](), false
	}
	removingEl := l.arr[idx]

	l.shiftLeft(idx)

	l.length--
	return removingEl, true
}

func (l *ArrayList[T]) increaseArrSize() {
	l.cap *= 2
	newArr := make([]T, l.cap)

	for i := 0; i < l.length; i++ {
		newArr[i] = l.arr[i]
	}

	l.arr = newArr
}

func (l *ArrayList[T]) shiftLeft(from int) {
	for i := from; i < l.length-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.arr[l.length-1] = dsa.ZeroValue[T]()
}

func (l *ArrayList[T]) shiftRight(from int) {
	for i := l.length; i > from; i-- {
		l.arr[i] = l.arr[i-1]
	}
}
