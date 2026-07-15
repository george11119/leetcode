package max_area_of_island

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	tt := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 1},
			},
			want: 6,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := maxAreaOfIsland(tc.grid)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
