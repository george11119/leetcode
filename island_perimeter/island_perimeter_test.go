package island_perimeter

import "testing"

func TestIslandPerimeter(t *testing.T) {
	tt := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 0},
				{1, 1, 1, 0},
				{0, 0, 1, 1},
			},
			18,
		},
		{
			[][]int{
				{1, 0},
			},
			4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := islandPerimeter(tc.grid)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
