package stack

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestStack_PopOrder(t *testing.T) {
	tcs := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  nil,
		},
		{
			name:  "single",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "multiple",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			s := MakeNew(tc.input...)
			var got []int
			for s.Size() > 0 {
				got = append(got, s.Pop())
			}
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestStack_Peek(t *testing.T) {
	s := MakeNew(3, 2, 1)
	f.Assert(t, s.Peek()).Eq(1, "peek should be first value")
	f.Assert(t, s.Peek()).Eq(1, "peek shouldn't change")
	got := s.Pop()
	f.Assert(t, got).Eq(1, "peek should match first pop")
	f.Assert(t, s.Peek()).Eq(2, "peek should reflect new first value")
	f.Assert(t, s.Peek()).Eq(s.Pop(), "peek should match pop")

}
