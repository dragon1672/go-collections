package binarytree

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestCompare(t *testing.T) {
	tcs := []struct {
		name string
		a    *Node[int]
		b    *Node[int]
		want bool
	}{
		{
			name: "nil match",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "a nil not match",
			a:    nil,
			b:    &Node[int]{},
			want: false,
		},
		{
			name: "b nil not match",
			a:    &Node[int]{},
			b:    nil,
			want: false,
		},
		{
			name: "single match",
			a:    &Node[int]{Data: 1},
			b:    &Node[int]{Data: 1},
			want: true,
		},
		{
			name: "single don't match",
			a:    &Node[int]{Data: 1},
			b:    &Node[int]{Data: 2},
			want: false,
		},
		{
			name: "single child left",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}},
			want: true,
		},
		{
			name: "single child left different values",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Left: &Node[int]{Data: 2}},
			want: false,
		},
		{
			name: "single child right",
			a:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			want: true,
		},
		{
			name: "single child right different values",
			a:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Right: &Node[int]{Data: 2}},
			want: false,
		},
		{
			name: "single mix child a left b right",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			want: false,
		},
		{
			name: "single mix child a right b left",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			want: false,
		},
		{
			name: "single child right",
			a:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			b:    &Node[int]{Data: 0, Right: &Node[int]{Data: 1}},
			want: true,
		},
		{
			name: "both children match",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}, Right: &Node[int]{Data: 2}},
			b:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}, Right: &Node[int]{Data: 2}},
			want: true,
		},
		{
			name: "both children different values at leafs",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}, Right: &Node[int]{Data: 2}},
			b:    &Node[int]{Data: 0, Left: &Node[int]{Data: 3}, Right: &Node[int]{Data: 3}},
			want: false,
		},
		{
			name: "both children different root values",
			a:    &Node[int]{Data: 0, Left: &Node[int]{Data: 1}, Right: &Node[int]{Data: 2}},
			b:    &Node[int]{Data: 9, Left: &Node[int]{Data: 1}, Right: &Node[int]{Data: 2}},
			want: false,
		},
		{
			name: "large tree match",
			a: &Node[int]{
				Data: 0,
				Left: &Node[int]{
					Data:  1,
					Left:  &Node[int]{Data: 3},
					Right: &Node[int]{Data: 4},
				},
				Right: &Node[int]{
					Data:  2,
					Left:  &Node[int]{Data: 5},
					Right: &Node[int]{Data: 6},
				}},
			b: &Node[int]{
				Data: 0,
				Left: &Node[int]{
					Data:  1,
					Left:  &Node[int]{Data: 3},
					Right: &Node[int]{Data: 4},
				},
				Right: &Node[int]{
					Data:  2,
					Left:  &Node[int]{Data: 5},
					Right: &Node[int]{Data: 6},
				}},
			want: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := BinaryTreesSame(tc.a, tc.b)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
