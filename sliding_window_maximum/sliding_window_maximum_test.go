package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 1, 0, 4, 2, 6},
			3,
			[]int{2, 2, 4, 4, 6},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1, -1},
			1,
			[]int{1, -1},
		},
		{
			[]int{1, 3, 1, 2, 0, 5},
			3,
			[]int{3, 3, 2, 5},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := maxSlidingWindow(tc.nums, tc.k)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
