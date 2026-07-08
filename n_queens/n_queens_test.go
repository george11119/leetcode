package n_queens

import (
	"fmt"
	"reflect"
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

func TestIsSafe(t *testing.T) {
	tt := []struct {
		board [][]bool
		row   int
		col   int
		want  bool
	}{
		{
			[][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{true, false, false, false},
			},
			0,
			3,
			false,
		},
		{
			[][]bool{
				{false, false, false, true},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			3,
			0,
			false,
		},
		{
			[][]bool{
				{true, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			3,
			3,
			false,
		},
		{
			[][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, true},
			},
			0,
			0,
			false,
		},
		{
			[][]bool{
				{true, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			3,
			1,
			true,
		},
		{
			[][]bool{
				{true, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			3,
			0,
			false,
		},
		{
			[][]bool{
				{true, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			0,
			3,
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := isSafe(tc.board, tc.row, tc.col)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}

func TestConvertToStringSlice(t *testing.T) {
	tt := []struct {
		board [][]bool
		want  []string
	}{
		{
			[][]bool{
				{false, true, false, false},
				{false, false, false, true},
				{true, false, false, false},
				{false, false, true, false},
			},
			[]string{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := convertToStringSlice(tc.board)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}

func TestNQueens(t *testing.T) {
	tt := []struct {
		n    int
		want [][]string
	}{
		{
			4,
			[][]string{
				{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.",
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := solveNQueens(tc.n)
			want := tc.want

			if !equalUnordered(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
