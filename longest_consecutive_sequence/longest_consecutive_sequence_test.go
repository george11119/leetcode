package longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tt := []struct {
		input []int
		want  int
	}{
		{
			[]int{2, 20, 4, 10, 3, 4, 5},
			4, // 2,3,4,5
		},
		{
			[]int{0, 3, 2, 5, 4, 6, 1, 1},
			7, // 0,1,2,3,4,5,6
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := longestConsecutive(tc.input)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
