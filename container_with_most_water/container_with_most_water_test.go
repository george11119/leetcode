package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	tt := []struct {
		heights []int
		want    int
	}{
		{
			[]int{1, 7, 2, 5, 4, 7, 3, 6},
			36,
		},
		{
			[]int{2, 2, 2},
			4,
		},
		{
			[]int{5, 10, 30, 30, 10, 1},
			30,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := maxArea(tc.heights)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
