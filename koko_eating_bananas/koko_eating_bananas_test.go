package koko_eating_bananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tt := []struct {
		piles []int
		h     int
		want  int
	}{
		{
			[]int{1, 4, 3, 2},
			9,
			2,
		},
		{
			[]int{25, 10, 23, 4},
			4,
			25,
		},
		{
			[]int{1, 4, 3, 2, 7},
			9,
			3,
		},
		{
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			[]int{
				332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184,
			},
			823855818,
			14,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := minEatingSpeed(tc.piles, tc.h)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
