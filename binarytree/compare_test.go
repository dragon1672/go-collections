package binarytree

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestCompare(t *testing.T) {
	tcs := []struct {
		name string
		a    *BinaryNode[int]
		b    *BinaryNode[int]
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
			b:    &BinaryNode[int]{},
			want: false,
		},
		{
			name: "b nil not match",
			a:    &BinaryNode[int]{},
			b:    nil,
			want: false,
		},
		{
			name: "single match",
			a:    &BinaryNode[int]{Data: 1},
			b:    &BinaryNode[int]{Data: 1},
			want: true,
		},
		{
			name: "single don't match",
			a:    &BinaryNode[int]{Data: 1},
			b:    &BinaryNode[int]{Data: 2},
			want: false,
		},
		{
			name: "single child left",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}},
			want: true,
		},
		{
			name: "single child left different values",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 2}},
			want: false,
		},
		{
			name: "single child right",
			a:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			want: true,
		},
		{
			name: "single child right different values",
			a:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 2}},
			want: false,
		},
		{
			name: "single mix child a left b right",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			want: false,
		},
		{
			name: "single mix child a right b left",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			want: false,
		},
		{
			name: "single child right",
			a:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			b:    &BinaryNode[int]{Data: 0, Right: &BinaryNode[int]{Data: 1}},
			want: true,
		},
		{
			name: "both children match",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}, Right: &BinaryNode[int]{Data: 2}},
			b:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}, Right: &BinaryNode[int]{Data: 2}},
			want: true,
		},
		{
			name: "both children different values at leafs",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}, Right: &BinaryNode[int]{Data: 2}},
			b:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 3}, Right: &BinaryNode[int]{Data: 3}},
			want: false,
		},
		{
			name: "both children different root values",
			a:    &BinaryNode[int]{Data: 0, Left: &BinaryNode[int]{Data: 1}, Right: &BinaryNode[int]{Data: 2}},
			b:    &BinaryNode[int]{Data: 9, Left: &BinaryNode[int]{Data: 1}, Right: &BinaryNode[int]{Data: 2}},
			want: false,
		},
		{
			name: "large tree match",
			a: &BinaryNode[int]{
				Data: 0,
				Left: &BinaryNode[int]{
					Data:  1,
					Left:  &BinaryNode[int]{Data: 3},
					Right: &BinaryNode[int]{Data: 4},
				},
				Right: &BinaryNode[int]{
					Data:  2,
					Left:  &BinaryNode[int]{Data: 5},
					Right: &BinaryNode[int]{Data: 6},
				}},
			b: &BinaryNode[int]{
				Data: 0,
				Left: &BinaryNode[int]{
					Data:  1,
					Left:  &BinaryNode[int]{Data: 3},
					Right: &BinaryNode[int]{Data: 4},
				},
				Right: &BinaryNode[int]{
					Data:  2,
					Left:  &BinaryNode[int]{Data: 5},
					Right: &BinaryNode[int]{Data: 6},
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
