package binaryprinter

import (
	"github.com/dragon1672/go-collections/binarytree"
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestWriteBinaryTree(t *testing.T) {
	tcs := []struct {
		name      string
		formatter *TrimmableBoard[string]
		input     *binarytree.Node[string]
		want      string
		wantErr   bool
	}{
		{
			name:      "nil",
			formatter: &TrimmableBoard[string]{},
			input:     nil,
			want:      "",
		},
		{
			name:      "single value",
			formatter: &TrimmableBoard[string]{},
			input:     &binarytree.Node[string]{Data: "hello there"},
			want:      "hello there",
		},
		{
			name:      "single leve of children",
			formatter: &TrimmableBoard[string]{},
			input: &binarytree.Node[string]{
				Data:  "root",
				Left:  &binarytree.Node[string]{Data: "L1"},
				Right: &binarytree.Node[string]{Data: "R1"},
			},
			want: "" +
				" /root \\\n" +
				"L1    R1",
		},
		{
			name:      "single leve of children with padding",
			formatter: &TrimmableBoard[string]{NodeSpacing: 3},
			input: &binarytree.Node[string]{
				Data:  "root",
				Left:  &binarytree.Node[string]{Data: "L1"},
				Right: &binarytree.Node[string]{Data: "R1"},
			},
			want: "" +
				" /root-\\\n" +
				"L1    R1",
		},
		{
			name:      "single child left",
			formatter: &TrimmableBoard[string]{},
			input: &binarytree.Node[string]{
				Data: "root",
				Left: &binarytree.Node[string]{Data: "L1"},
			},
			want: "" +
				" /root\n" +
				"L1    ",
		},
		{
			name:      "single child right",
			formatter: &TrimmableBoard[string]{},
			input: &binarytree.Node[string]{
				Data:  "root",
				Right: &binarytree.Node[string]{Data: "R1"},
			},
			want: "" +
				"root \\\n" +
				"    R1",
		},
		{
			name:      "balanced tree",
			formatter: &TrimmableBoard[string]{},
			input: &binarytree.Node[string]{
				Data: "0",
				Left: &binarytree.Node[string]{
					Data: "1",
					Left: &binarytree.Node[string]{
						Data: "2",
					},
					Right: &binarytree.Node[string]{
						Data: "3",
					},
				},
				Right: &binarytree.Node[string]{
					Data: "4",
					Left: &binarytree.Node[string]{
						Data: "5",
					},
					Right: &binarytree.Node[string]{
						Data: "6",
					},
				},
			},
			want: "" +
				" /-0-\\ \n" +
				"/1\\ /4\\\n" +
				"2 3 5 6",
		},
		{
			name: "balanced tree padding no collapse",
			formatter: &TrimmableBoard[string]{
				NoCollapse:  true,
				NodeSpacing: 3,
			},
			input: &binarytree.Node[string]{
				Data: "0",
				Left: &binarytree.Node[string]{
					Data: "1",
					Left: &binarytree.Node[string]{
						Data: "2",
					},
					Right: &binarytree.Node[string]{
						Data: "3",
					},
				},
				Right: &binarytree.Node[string]{
					Data: "4",
					Left: &binarytree.Node[string]{
						Data: "5",
					},
					Right: &binarytree.Node[string]{
						Data: "6",
					},
				},
			},
			want: "" +
				"   /---0---\\  \n" +
				" /-1-\\   /-4-\\\n" +
				" 2   3   5   6",
		},
		{
			name:      "unbalanced tree",
			formatter: &TrimmableBoard[string]{},
			input: &binarytree.Node[string]{
				Data: "0",
				Left: &binarytree.Node[string]{
					Data: "1",
					Left: &binarytree.Node[string]{
						Data: "3",
						Left: &binarytree.Node[string]{
							Data: "4",
						},
					},
					Right: &binarytree.Node[string]{
						Data: "2",
					},
				},
			},
			want: "" +
				"  /-0\n" +
				" /1\\ \n" +
				"/3 2 \n" +
				"4    ",
		},
		{
			name:      "unbalanced tree uncollapsed",
			formatter: &TrimmableBoard[string]{NoCollapse: true},
			input: &binarytree.Node[string]{
				Data: "0",
				Left: &binarytree.Node[string]{
					Data: "1",
					Left: &binarytree.Node[string]{
						Data: "3",
						Left: &binarytree.Node[string]{
							Data: "4",
						},
					},
					Right: &binarytree.Node[string]{
						Data: "2",
					},
				},
			},
			want: "" +
				"   /---0\n" +
				" /-1-\\  \n" +
				"/3   2  \n" +
				"4       ",
		},
		{
			name:      "huffman example",
			formatter: &TrimmableBoard[string]{NoCollapse: true},
			input: &binarytree.Node[string]{
				Data: "0",
				Left: &binarytree.Node[string]{
					Data: "nul(13)",
					Left: &binarytree.Node[string]{
						Data: "nul(5)",
						Left: &binarytree.Node[string]{
							Data: "67(2)",
						},
						Right: &binarytree.Node[string]{
							Data: "nul(3)",
							Left: &binarytree.Node[string]{
								Data: "34(1)",
							},
							Right: &binarytree.Node[string]{
								Data: "nul(2)",
								Left: &binarytree.Node[string]{
									Data: "78(1)",
								},
								Right: &binarytree.Node[string]{
									Data: "89(1)",
								},
							},
						},
					},
					Right: &binarytree.Node[string]{
						Data: "nul(8)",
						Left: &binarytree.Node[string]{
							Data: "45(4)",
						},
						Right: &binarytree.Node[string]{
							Data: "nul(4)",
							Left: &binarytree.Node[string]{
								Data: "23(2)",
							},
							Right: &binarytree.Node[string]{
								Data: "nul(2)",
								Left: &binarytree.Node[string]{
									Data: "12(1)",
								},
								Right: &binarytree.Node[string]{
									Data: "56(1)",
								},
							},
						},
					},
				},
			},
			want: "" +
				"                                                               /------------------------------------------------------------   0\n" +
				"                               /----------------------------nul(13)----------------------------\\                                \n" +
				"               /------------nul(5) ------------\\                               /------------nul(8) ------------\\                \n" +
				"             67(2)                     /----nul(3) ----\\                     45(4)                     /----nul(4) ----\\        \n" +
				"                                     34(1)         /nul(2) \\                                         23(2)         /nul(2) \\    \n" +
				"                                                 78(1)   89(1)                                                   12(1)   56(1)  ",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			gotErr := tc.formatter.WriteBinaryTree(tc.input)
			f.Assert(t, gotErr != nil).Eq(tc.wantErr, "expected error")
			got := tc.formatter.String()
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
