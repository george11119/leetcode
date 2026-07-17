package rotting_fruit

import "testing"

func TestOrangesRotting(t *testing.T) {
	tt := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 2},
			},
			4,
		},
		{
			[][]int{
				{1, 0, 1},
				{0, 2, 0},
				{1, 0, 1},
			},
			-1,
		},
		{
			[][]int{
				{2, 1, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			-1,
		},
		{
			[][]int{
				{2, 2},
				{1, 1},
				{0, 0},
				{2, 0},
			},
			1,
		},
		{
			[][]int{{0}},
			0,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := orangesRotting(tc.grid)
			want := tc.want
			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
