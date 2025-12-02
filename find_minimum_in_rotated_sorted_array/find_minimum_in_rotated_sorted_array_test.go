package find_minimum_in_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			[]int{3, 4, 5, 6, 1, 2},
			1,
		},
		{
			[]int{4, 5, 0, 1, 2, 3},
			0,
		},
		{
			[]int{4, 5, 6, 7},
			4,
		},
		{
			[]int{2, 3, 4, 5, 6, 1},
			1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findMin(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
