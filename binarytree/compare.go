package binarytree

import (
	"github.com/dragon1672/go-collections/queue"
)

type treeCompare[T any] struct {
	a, b *Node[T]
}

func BinaryTreesSame[T comparable](a, b *Node[T]) bool {
	traverse := queue.MakeNew(treeCompare[T]{a: a, b: b})
	seen := make(map[*Node[T]]bool)
	for traverse.Size() > 0 {
		elem := traverse.Pop()
		if _, exists := seen[elem.a]; exists {
			continue
		}
		seen[elem.a] = true // track with only one tree

		if elem.a == nil && elem.b == nil {
			continue
		}
		if elem.a == nil || elem.b == nil {
			return false
		}
		if elem.a.Data != elem.b.Data {
			return false
		}

		// enqueue the children
		traverse.Push(treeCompare[T]{a: elem.a.Right, b: elem.b.Right})
		traverse.Push(treeCompare[T]{a: elem.a.Left, b: elem.b.Left})
	}
	return true
}
