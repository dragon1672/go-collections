package binarytree

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestHeight(t *testing.T) {
	tcs := []struct {
		name    string
		n       *BinaryNode[string]
		want    int
		wantErr bool
	}{
		{
			name: "nil",
			n:    nil,
			want: 0,
		},
		{
			name: "single",
			n:    &BinaryNode[string]{},
			want: 1,
		},
		{
			name: "single child left",
			n:    &BinaryNode[string]{Left: &BinaryNode[string]{}},
			want: 2,
		},
		{
			name: "single child right",
			n:    &BinaryNode[string]{Right: &BinaryNode[string]{}},
			want: 2,
		},
		{
			name: "both children",
			n:    &BinaryNode[string]{Left: &BinaryNode[string]{}, Right: &BinaryNode[string]{}},
			want: 2,
		},
		{
			name: "big balanced tree",
			n: &BinaryNode[string]{
				Left: &BinaryNode[string]{
					Left: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
						Right: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
					},
					Right: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
						Right: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
					},
				},
				Right: &BinaryNode[string]{
					Left: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
						Right: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
					},
					Right: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
						Right: &BinaryNode[string]{
							Left:  &BinaryNode[string]{},
							Right: &BinaryNode[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big unbalanced left tree",
			n: &BinaryNode[string]{
				Left: &BinaryNode[string]{
					Left: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Left: &BinaryNode[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big unbalanced right tree",
			n: &BinaryNode[string]{
				Right: &BinaryNode[string]{
					Right: &BinaryNode[string]{
						Right: &BinaryNode[string]{
							Right: &BinaryNode[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big zig zag tree",
			n: &BinaryNode[string]{
				Left: &BinaryNode[string]{
					Right: &BinaryNode[string]{
						Left: &BinaryNode[string]{
							Right: &BinaryNode[string]{},
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := Height(tc.n)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
			f.Assert(t, gotErr != nil).Eq(tc.wantErr, "expected error")
		})
	}
}
