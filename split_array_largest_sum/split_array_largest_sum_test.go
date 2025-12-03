package split_array_largest_sum

import "testing"

func TestSplitArray(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want int
	}{
		{
			[]int{2, 4, 10, 1, 5},
			2,
			16,
		},
		{
			[]int{1, 0, 2, 3, 5},
			4,
			5,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := splitArray(tc.nums, tc.k)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
