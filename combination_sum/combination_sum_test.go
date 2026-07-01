package combination_sum

import (
	"fmt"
	"slices"
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
		slices.Sort(s)
		counts[fmt.Sprint(s)]++
	}
	for _, s := range want {
		slices.Sort(s)
		counts[fmt.Sprint(s)]--
	}

	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestCombinationSum(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			[]int{2, 5, 6, 9},
			9,
			[][]int{{2, 2, 5}, {9}},
		},
		{
			[]int{3, 4, 5},
			16,
			[][]int{{3, 3, 3, 3, 4}, {3, 3, 5, 5}, {4, 4, 4, 4}, {3, 4, 4, 5}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := combinationSum(tc.nums, tc.target)
			want := tc.want
			if !equalUnordered(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
