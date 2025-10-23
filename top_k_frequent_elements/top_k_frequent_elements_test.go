package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 2, 3, 3, 3}, 2, []int{3, 2},
		},
		{
			[]int{7, 7}, 1, []int{7},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := topKFrequent(tc.nums, tc.k)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
