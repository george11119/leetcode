package number_of_islands

import "testing"

func TestNumIslands(t *testing.T) {
	tt := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				{'0', '1', '1', '1', '0'},
				{'0', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '1'},
				{'1', '1', '0', '0', '1'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := numIslands(tc.grid)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
