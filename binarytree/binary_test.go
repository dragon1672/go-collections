package binarytree

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestHeight(t *testing.T) {
	tcs := []struct {
		name    string
		n       *Node[string]
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
			n:    &Node[string]{},
			want: 1,
		},
		{
			name: "single child left",
			n:    &Node[string]{Left: &Node[string]{}},
			want: 2,
		},
		{
			name: "single child right",
			n:    &Node[string]{Right: &Node[string]{}},
			want: 2,
		},
		{
			name: "both children",
			n:    &Node[string]{Left: &Node[string]{}, Right: &Node[string]{}},
			want: 2,
		},
		{
			name: "big balanced tree",
			n: &Node[string]{
				Left: &Node[string]{
					Left: &Node[string]{
						Left: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
						Right: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
					},
					Right: &Node[string]{
						Left: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
						Right: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
					},
				},
				Right: &Node[string]{
					Left: &Node[string]{
						Left: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
						Right: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
					},
					Right: &Node[string]{
						Left: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
						Right: &Node[string]{
							Left:  &Node[string]{},
							Right: &Node[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big unbalanced left tree",
			n: &Node[string]{
				Left: &Node[string]{
					Left: &Node[string]{
						Left: &Node[string]{
							Left: &Node[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big unbalanced right tree",
			n: &Node[string]{
				Right: &Node[string]{
					Right: &Node[string]{
						Right: &Node[string]{
							Right: &Node[string]{},
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "big zig zag tree",
			n: &Node[string]{
				Left: &Node[string]{
					Right: &Node[string]{
						Left: &Node[string]{
							Right: &Node[string]{},
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
