package combination_sum_ii

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
			[]int{9, 2, 2, 4, 6, 1, 5},
			8,
			[][]int{{1, 2, 5}, {2, 2, 4}, {2, 6}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			7,
			[][]int{{1, 2, 4}, {2, 5}, {3, 4}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := combinationSum2(tc.nums, tc.target)
			want := tc.want
			if !equalUnordered(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
