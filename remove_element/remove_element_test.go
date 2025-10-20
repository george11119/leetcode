package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tt := []struct {
		nums []int
		val  int
		k    int
		want []int
	}{
		{
			[]int{1, 1, 2, 3, 4},
			1,
			3, // Explanation: You should return k = 3 as we have 3 elements which are not equal to val = 1.
			[]int{2, 3, 4, 1, 1},
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5, // Explanation: You should return k = 5 as we have 5 elements which are not equal to val = 2.
			[]int{0, 1, 3, 0, 4, 2, 2, 2},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := removeElement(&tc.nums, tc.val)
			//removeElement(tc.nums, tc.val)

			if got != tc.k {
				t.Fatalf("got %v, k %v", got, tc.k)
			}

			for i := 0; i < tc.k; i++ {
				if tc.nums[i] != tc.want[i] {
					t.Fatalf("nums %v, want %v", tc.nums[0:tc.k], tc.want[0:tc.k])
				}
			}
		})
	}
}
