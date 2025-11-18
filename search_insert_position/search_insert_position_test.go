package search_insert_position

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			[]int{-1, 0, 2, 4, 6, 8},
			5,
			4,
		},
		{
			[]int{-1, 0, 2, 4, 6, 8},
			10,
			6,
		},
		{
			[]int{-1, 0, 2, 4, 6, 8},
			-2,
			0,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
