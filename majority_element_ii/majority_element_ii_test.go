package majority_element_ii

import (
	"reflect"
	"sort"
	"testing"
)

func compareArrays[T any](a, b []T) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	return reflect.DeepEqual(a, b)
}

func TestMajorityElement(t *testing.T) {
	tt := []struct {
		input []int
		want  []int
	}{
		{
			[]int{5, 2, 3, 2, 2, 2, 2, 5, 5, 5},
			[]int{2, 5},
		},
		{
			[]int{4, 4, 4, 4, 4},
			[]int{4},
		},
		{
			[]int{1, 2, 3},
			[]int{},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := majorityElement(tc.input)
			want := tc.want

			// order of values returned from got doesn't matter. Sort result to make test not flaky
			sort.Ints(got)
			sort.Ints(want)

			if !compareArrays[int](got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
