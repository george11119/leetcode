package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tt := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 4, 6},
			[]int{48, 24, 12, 8},
		},
		{
			[]int{-1, 0, 1, 2, 3},
			[]int{0, -6, 0, 0, 0},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := ProductExceptSelf(tc.input)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
