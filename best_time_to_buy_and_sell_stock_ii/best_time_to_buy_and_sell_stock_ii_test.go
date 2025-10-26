package best_time_to_buy_and_sell_stock_ii

import "testing"

func TestMaxProfit(t *testing.T) {
	tt := []struct {
		input []int
		want  int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := maxProfit(tc.input)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
