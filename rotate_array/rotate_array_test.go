package rotate_array

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			4,
			[]int{5, 6, 7, 8, 1, 2, 3, 4},
		},
		{
			[]int{1000, 2, 4, -3},
			2,
			[]int{4, -3, 1000, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			7,
			[]int{4, 5, 1, 2, 3},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			rotate(tc.nums, tc.k)
			if !reflect.DeepEqual(tc.want, tc.nums) {
				t.Fatalf("got %v, want %v", tc.nums, tc.want)
			}
		})
	}
}
