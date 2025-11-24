package capacity_to_ship_within_d_days

import "testing"

func TestShipWithinDays(t *testing.T) {
	tt := []struct {
		weights []int
		days    int
		want    int
	}{
		{
			[]int{2, 4, 6, 1, 3, 10},
			4,
			10,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			5,
		},
		{
			[]int{1, 5, 4, 4, 2, 3},
			3,
			8,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := shipWithinDays(tc.weights, tc.days)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
