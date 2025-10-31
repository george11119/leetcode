package two_integer_sum_ii

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			[]int{1, 2, 3, 4},
			3,
			[]int{1, 2},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
		{
			[]int{-5, -3, 0, 2, 4, 6, 8},
			5,
			[]int{2, 7},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := twoSum(tc.numbers, tc.target)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
