package palindrome_partitioning

import (
	"fmt"
	"testing"
)

// equalUnordered reports whether got and want contain the same inner slices.
// Each inner slice must match exactly (same elements in the same order), but
// the inner slices may appear in any order within the outer slice.
func equalUnordered(got, want [][]string) bool {
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

func TestPalindromePartitioning(t *testing.T) {
	tt := []struct {
		s    string
		want [][]string
	}{
		{
			"a",
			[][]string{{"a"}},
		},
		{
			"aab",
			[][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := partition(tc.s)
			want := tc.want

			if !equalUnordered(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
