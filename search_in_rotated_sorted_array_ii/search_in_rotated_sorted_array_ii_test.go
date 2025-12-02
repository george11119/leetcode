package search_in_rotated_sorted_array_ii

import "testing"

func TestSearch(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   bool
	}{
		{
			[]int{3, 4, 4, 5, 6, 1, 2, 2},
			1,
			true,
		},
		{
			[]int{3, 5, 6, 0, 0, 1, 2},
			4,
			false,
		},
		{
			[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			7,
			true,
		},
		{
			[]int{1, 1, 3, 1},
			3,
			true,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := search(tc.nums, tc.target)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
