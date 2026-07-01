package subsets

import (
	"fmt"
	"testing"
)

// equalUnordered reports whether got and want contain the same inner slices.
// Each inner slice must match exactly (same elements in the same order), but
// the inner slices may appear in any order within the outer slice.
func equalUnordered(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	counts := make(map[string]int)
	for _, s := range got {
		counts[fmt.Sprint(s)]++
	}
	for _, s := range want {
		counts[fmt.Sprint(s)]--
	}

	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestSubsets(t *testing.T) {
	tt := []struct {
		nums []int
		want [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := subsets(tc.nums)
			want := tc.want
			if !equalUnordered(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
