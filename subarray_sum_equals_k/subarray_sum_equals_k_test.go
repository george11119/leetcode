package subarray_sum_equals_k

import "testing"

func TestSubarraySum(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want int
	}{
		{
			[]int{2, -1, 1, 2},
			2,
			4,
		},
		{
			[]int{4, 4, 4, 4, 4, 4},
			4,
			6,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := subarraySum(tc.nums, tc.k)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
