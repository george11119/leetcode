package pacific_atlantic_water_flow

import (
	"cmp"
	"slices"
	"testing"
)

func EqualUnordered[T cmp.Ordered](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	ac := slices.Clone(a)
	bc := slices.Clone(b)
	slices.SortFunc(ac, slices.Compare[[]T])
	slices.SortFunc(bc, slices.Compare[[]T])
	return slices.EqualFunc(ac, bc, slices.Equal[[]T])
}

func TestPacificAtlantic(t *testing.T) {
	tt := []struct {
		heights [][]int
		want    [][]int
	}{
		{
			[][]int{
				{4, 2, 6, 3, 4},
				{7, 4, 6, 4, 7},
				{6, 3, 5, 3, 6},
			},
			[][]int{
				{0, 2}, {0, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 0},
			},
		},
		{
			[][]int{
				{1},
			},
			[][]int{{0, 0}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := pacificAtlantic(tc.heights)
			want := tc.want

			if !EqualUnordered(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
