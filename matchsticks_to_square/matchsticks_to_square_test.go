package matchsticks_to_square

import "testing"

func TestMatchsticksToSquare(t *testing.T) {
	tt := []struct {
		matchsticks []int
		want        bool
	}{
		{
			[]int{1, 3, 4, 2, 2, 4},
			true,
		},
		{
			[]int{1, 1, 2, 2, 2},
			true,
		},
		{
			[]int{1, 5, 6, 3},
			false,
		},
		{
			[]int{3, 3, 3, 3, 4},
			false,
		},
		{
			[]int{15, 15, 15, 14, 12, 12, 12, 12, 12, 11, 10, 9, 5, 1, 1},
			true,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := makesquare(tc.matchsticks)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
