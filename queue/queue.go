// Package queue is a FIFO queue.
package queue

type node[T any] struct {
	data T
	next *node[T]
}

// Queue is a FIFO queue.
type Queue[T any] struct {
	head, tail *node[T]
	count      int
}

func (q *Queue[T]) Push(e T) {
	wrapped := &node[T]{data: e}
	if q.tail == nil {
		q.head = wrapped
		q.tail = wrapped
	} else {
		q.tail.next = wrapped
		q.tail = wrapped
	}
	q.count++
}

func (q *Queue[T]) Pop() T {
	if q.head == nil {
		var empty T
		return empty
	}
	ret := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.count--
	return ret
}

func (q *Queue[T]) Peek() T {
	if q.head == nil {
		var empty T
		return empty
	}
	return q.head.data
}

func (q *Queue[T]) Len() int {
	return q.count
}

func FromSlice[T any](s []T) *Queue[T] {
	ret := &Queue[T]{}
	for _, e := range s {
		ret.Push(e)
	}
	return ret
}
