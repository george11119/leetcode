package find_the_duplicate_number

import "testing"

func TestFindDuplicate(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 2, 3, 2, 2},
			2,
		},
		{
			[]int{1, 2, 3, 4, 4},
			4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findDuplicate(tc.nums)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
