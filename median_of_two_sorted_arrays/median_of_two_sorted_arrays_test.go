package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tt := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{ // small cases
			[]int{1, 2},
			[]int{3},
			2.0,
		},
		{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			[]int{1, 3},
			[]int{2, 4},
			2.5,
		},
		{ // case where we check left
			[]int{1, 3, 8, 9, 15},
			[]int{7, 11, 18, 19, 21, 25},
			11.0,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findMedianSortedArrays(tc.nums1, tc.nums2)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
