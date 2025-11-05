package minimum_size_subarray_sum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tt := []struct {
		target int
		nums   []int
		want   int
	}{
		{
			10,
			[]int{2, 1, 5, 1, 5, 3},
			3,
		},
		{
			5,
			[]int{1, 2, 1},
			0,
		},
		{
			7,
			[]int{2, 3, 1, 2, 4, 3},
			2,
		},
		{
			4,
			[]int{1, 4, 4},
			1,
		},
		{
			15,
			[]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			2,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := minSubArrayLen(tc.target, tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
