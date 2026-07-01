package combinations

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
		n    int
		k    int
		want [][]int
	}{
		{
			3,
			2,
			[][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			3,
			3,
			[][]int{{1, 2, 3}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := combine(tc.n, tc.k)
			want := tc.want
			if !equalUnordered(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
