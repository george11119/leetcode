package number_of_connected_components_in_an_undirected_graph

import "testing"

func TestCountComponents(t *testing.T) {
	tt := []struct {
		n     int
		edges [][]int
		want  int
	}{
		{
			5,
			[][]int{{0, 1}, {1, 2}, {3, 4}},
			2,
		},
		{
			5,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := countComponents(tc.n, tc.edges)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
