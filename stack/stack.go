// Package stack is a LIFO stack.
package stack

// Stack is a LIFO stack.
type Stack[T any] struct {
	elems []T
	count int
}

func (s *Stack[T]) Push(n T) {
	s.elems = append(s.elems[:s.count], n)
	s.count++
}

func (s *Stack[T]) Pop() T {
	if s.count <= 0 {
		var empty T
		return empty
	}
	s.count--
	return s.elems[s.count]
}

func (s *Stack[T]) Peek() T {
	if s.count <= 0 {
		var empty T
		return empty
	}
	return s.elems[s.count-1]
}

func (s *Stack[T]) Len() int {
	return s.count
}

func FromSlice[T any](s []T) *Stack[T] {
	ret := &Stack[T]{}
	for _, e := range s {
		ret.Push(e)
	}
	return ret
}
