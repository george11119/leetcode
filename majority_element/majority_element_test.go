package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			[]int{5, 5, 1, 1, 1, 5, 5},
			5,
		},
		{
			[]int{2, 2, 2},
			2,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := majorityElement(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
