package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"case 1", []int{3, 4, 5, 6}, 7, []int{0, 1}},
		{"case 2", []int{4, 5, 6}, 10, []int{0, 2}},
		{"case 3", []int{5, 5}, 10, []int{0, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			want := tc.want

			if reflect.DeepEqual(got, want) != true {
				t.Fatalf("want: %v, got: %v", want, got)
			}
		})
	}
}
