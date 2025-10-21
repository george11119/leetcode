package merge_sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tt := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			[]int{2, 5},
			[]int{1, 3},
			[]int{1, 2, 3, 5},
		},
		{
			[]int{2, 5, 7},
			[]int{1, 3},
			[]int{1, 2, 3, 5, 7},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := Merge(tc.arr1, tc.arr2)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tt := []struct {
		inputArr []int
		want     []int
	}{
		{
			[]int{7, 2, 1, 8, 6, 3, 5, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{7, 2, 1, 8, 9, 6, 3, 5, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := MergeSort(tc.inputArr)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
