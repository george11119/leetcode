package partition_to_k_equal_subsets

import "testing"

func TestCanPartitionKSubsets(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			[]int{2, 4, 1, 3, 5},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 4},
			3,
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			8,
			true,
		},
		{
			[]int{1, 1, 2, 2, 2},
			3,
			false,
		},
		{
			[]int{4, 5, 9, 3, 10, 2, 10, 7, 10, 8, 5, 9, 4, 6, 4, 9},
			5,
			true,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := canPartitionKSubsets(tc.nums, tc.k)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
