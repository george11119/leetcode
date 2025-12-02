package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		//{
		//	[]int{3, 4, 5, 6, 1, 2},
		//	1,
		//	4,
		//},
		//{
		//	[]int{3, 5, 6, 0, 1, 2},
		//	4,
		//	-1,
		//},
		{
			[]int{5, 1, 3},
			3,
			2,
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
