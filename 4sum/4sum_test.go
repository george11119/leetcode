package _4sum

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

func TestFourSum(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			[]int{3, 2, 3, -3, 1, 0},
			3,
			[][]int{{-3, 0, 3, 3}, {-3, 1, 2, 3}},
		},
		{
			[]int{1, -1, 1, -1, 1, -1},
			2,
			[][]int{{-1, 1, 1, 1}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4},
			-1,
			[][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}},
		},
		{
			[]int{2, 2, 2, 2, 2},
			8,
			[][]int{{2, 2, 2, 2}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := fourSum(tc.nums, tc.target)
			want := tc.want

			if !arrayEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
