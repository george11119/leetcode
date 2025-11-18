package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			[]int{-1, 0, 2, 4, 6, 8},
			4,
			3,
		},
		{
			[]int{-1, 0, 2, 4, 6, 8},
			3,
			-1,
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
