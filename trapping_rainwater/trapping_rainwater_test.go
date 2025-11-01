package trapping_rainwater

import "testing"

func TestTrap(t *testing.T) {
	tt := []struct {
		height []int
		want   int
	}{
		{
			[]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			9,
		},
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := trap(tc.height)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
