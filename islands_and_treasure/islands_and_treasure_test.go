package islands_and_treasure

import (
	"reflect"
	"testing"
)

func TestIslandsAndTreasure(t *testing.T) {
	tt := []struct {
		grid [][]int
		want [][]int
	}{
		{
			[][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			[][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := tc.grid
			islandsAndTreasure(got)
			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
