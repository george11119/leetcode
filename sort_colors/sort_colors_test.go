package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tt := []struct {
		inputArr []int
		want     []int
	}{
		{
			[]int{1, 0, 1, 2},
			[]int{0, 1, 1, 2},
		},
		{
			[]int{2, 1, 0},
			[]int{0, 1, 2},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			SortColors(tc.inputArr)

			if !reflect.DeepEqual(tc.inputArr, tc.want) {
				t.Fatalf("got %v, want %v", tc.inputArr, tc.want)
			}
		})
	}
}
