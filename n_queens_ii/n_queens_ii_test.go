package n_queens_ii

import (
	"testing"
)

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

func TestNQueens(t *testing.T) {
	tt := []struct {
		n    int
		want int
	}{
		{
			4,
			2,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := totalNQueens(tc.n)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
