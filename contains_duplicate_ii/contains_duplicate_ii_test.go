package contains_duplicate_ii

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tt := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			[]int{1, 2, 3, 1},
			3,
			true,
		},
		{
			[]int{2, 1, 2},
			1,
			false,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			2,
			false,
		},
		{
			[]int{1, 0, 1, 1},
			1,
			true,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := containsNearbyDuplicate(tc.nums, tc.k)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
