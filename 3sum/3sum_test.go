package _3sum

import (
	"reflect"
	"testing"
)

func arrayEqual(a, b [][]int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	return reflect.DeepEqual(a, b)
}

func TestThreeSum(t *testing.T) {
	tt := []struct {
		nums []int
		want [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{-2, 0, 0, 2, 2},
			[][]int{{-2, 0, 2}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := threeSum(tc.nums)
			want := tc.want

			if !arrayEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
