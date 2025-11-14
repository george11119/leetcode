package largest_rectangle_in_histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tt := []struct {
		heights []int
		want    int
	}{
		{
			[]int{7, 1, 7, 2, 2, 4},
			8,
		},
		{
			[]int{1, 3, 7},
			7,
		},
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			[]int{3, 6, 5, 7, 4, 8, 1, 0},
			20,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := largestRectangleArea(tc.heights)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
