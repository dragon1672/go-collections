package binarytree

import (
	"fmt"
	"github.com/dragon1672/go-collections/queue"
)

type BinaryNode[T any] struct {
	Data        T
	Left, Right *BinaryNode[T]
}

func (n *BinaryNode[T]) SetData(data T) *BinaryNode[T]            { n.Data = data; return n }
func (n *BinaryNode[T]) SetLeft(l *BinaryNode[T]) *BinaryNode[T]  { n.Left = l; return n }
func (n *BinaryNode[T]) SetRight(r *BinaryNode[T]) *BinaryNode[T] { n.Right = r; return n }

type depthMeter[T any] struct {
	depth int
	n     *BinaryNode[T]
}

func Height[T any](n *BinaryNode[T]) (int, error) {
	traverse := queue.MakeNew(depthMeter[T]{depth: 1, n: n})
	maxDepth := 0
	seen := make(map[*BinaryNode[T]]bool)
	for traverse.Size() > 0 {
		elem := traverse.Pop()
		if elem.n == nil {
			continue
		}
		if _, exists := seen[elem.n]; exists {
			return -1, fmt.Errorf("cycle detected, tree is not valid binary tree")
		}
		seen[elem.n] = true
		if elem.depth > maxDepth {
			maxDepth = elem.depth
		}

		traverse.Push(depthMeter[T]{
			depth: elem.depth + 1,
			n:     elem.n.Right,
		})
		traverse.Push(depthMeter[T]{
			depth: elem.depth + 1,
			n:     elem.n.Left,
		})
	}
	return maxDepth, nil
}
