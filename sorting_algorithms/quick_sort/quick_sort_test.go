package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tt := []struct {
		inputArr []int
		want     []int
	}{
		{
			[]int{7, 2, 1, 8, 6, 3, 5, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{7, 2, 1, 8, 9, 6, 3, 5, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := SortArray(tc.inputArr)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
