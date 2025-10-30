package merge_sorted_array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tt := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			[]int{10, 20, 20, 40, 0, 0},
			4,
			[]int{1, 2},
			2,
			[]int{1, 2, 10, 20, 20, 40},
		},
		{
			[]int{0, 0},
			0,
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 0, 0, 0, 0},
			2,
			[]int{10, 20, 20, 40},
			4,
			[]int{1, 2, 10, 20, 20, 40},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)

			if !reflect.DeepEqual(tc.nums1, tc.want) {
				t.Fatalf("got %v, want %v", tc.nums1, tc.want)
			}
		})
	}
}
