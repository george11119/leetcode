package find_k_closest_elements

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	tt := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			3,
			4,
			[]int{3, 4, 5},
		},
		{
			[]int{1, 3, 5, 7, 9, 11, 13},
			4,
			8,
			[]int{5, 7, 9, 11},
		},
		{
			[]int{2, 4, 5, 8},
			2,
			6,
			[]int{4, 5},
		},
		{
			[]int{2, 3, 4},
			3,
			1,
			[]int{2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
			10,
			[]int{2, 3, 4, 5},
		},
		{
			[]int{5, 10, 15, 20, 25},
			1,
			13,
			[]int{15},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findClosestElements(tc.arr, tc.k, tc.x)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
