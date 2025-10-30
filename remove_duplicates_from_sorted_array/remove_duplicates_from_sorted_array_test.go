package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		nums    []int
		want    int
		newNums []int
	}{
		{
			[]int{1, 1, 2, 3, 4},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{2, 10, 10, 30, 30, 30},
			3,
			[]int{2, 10, 30},
		},
		{
			[]int{-50, -50, -49, -48, -48, -47},
			4,
			[]int{-50, -49, -48, -47},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := removeDuplicates(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}

			for i := 0; i < len(tc.newNums); i++ {
				if tc.newNums[i] != tc.nums[i] {
					t.Fatalf("got %v want %v. nums = %v", tc.nums[:len(tc.newNums)], tc.newNums, tc.nums)
				}
			}
		})
	}
}
