package search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tt := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			10,
			true,
		},
		{
			[][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			15,
			false,
		},
		{
			[][]int{
				{1},
			},
			2,
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := searchMatrix(tc.matrix, tc.target)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
