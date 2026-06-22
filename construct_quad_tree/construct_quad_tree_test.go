package construct_quad_tree

import "testing"

func TestConstruct(t *testing.T) {
	tt := []struct {
		grid [][]int
	}{
		{
			[][]int{
				{0, 1, 1, 1, 0, 0, 0, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{0, 1, 1, 1, 0, 0, 0, 1},
			},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			construct(tc.grid)
		})
	}
}
