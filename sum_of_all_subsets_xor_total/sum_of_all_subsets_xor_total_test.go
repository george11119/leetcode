package sum_of_all_subsets_xor_total

import "testing"

func TestSumOfAllSubsetsXorTotal(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			[]int{2, 4, 6},
			24,
		},
		{
			[]int{2, 4},
			12,
		},
		{
			[]int{3, 1, 1},
			12,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := subsetXORSum(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
