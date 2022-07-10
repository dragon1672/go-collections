package binarytree

import (
	"fmt"
	"github.com/dragon1672/go-collections/queue"
)

type Node[T any] struct {
	Data        T
	Left, Right *Node[T]
}

type depthMeter[T any] struct {
	depth int
	n     *Node[T]
}

func Height[T any](n *Node[T]) (int, error) {
	traverse := queue.MakeNew(depthMeter[T]{depth: 1, n: n})
	maxDepth := 0
	seen := make(map[*Node[T]]bool)
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
