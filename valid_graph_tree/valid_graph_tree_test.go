package valid_graph_tree

import "testing"

func TestValidTree(t *testing.T) {
	tt := []struct {
		n     int
		edges [][]int
		want  bool
	}{
		{
			5,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			true,
		},
		{
			3,
			[][]int{{0, 1}, {0, 2}, {1, 2}},
			false,
		},
		{
			4,
			[][]int{{0, 1}, {2, 3}},
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := validTree(tc.n, tc.edges)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
