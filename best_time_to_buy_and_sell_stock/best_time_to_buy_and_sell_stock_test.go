package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tt := []struct {
		prices []int
		want   int
	}{
		{
			[]int{10, 1, 5, 6, 7, 1},
			6,
		},
		{
			[]int{10, 8, 7, 5, 2},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := maxProfit(tc.prices)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
