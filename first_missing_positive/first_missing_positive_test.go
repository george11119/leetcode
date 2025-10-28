package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			[]int{-2, -1, 0},
			1,
		},
		{
			[]int{1, 2, 4},
			3,
		},
		{
			[]int{1, 2, 3},
			4,
		},
		{
			[]int{1, 2, 4, 5, 6, 3, 1},
			7,
		},
		{
			[]int{3, -3, 6, 3},
			1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := firstMissingPositive(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
