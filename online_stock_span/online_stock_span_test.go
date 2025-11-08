package online_stock_span

import (
	"reflect"
	"testing"
)

func TestOnlineStockSpan(t *testing.T) {
	tt := []struct {
		nextVals []int
		want     []int
	}{
		{
			[]int{100, 80, 60, 70, 60, 75, 85},
			[]int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			[]int{100, 101, 102, 103},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{110, 110},
			[]int{1, 2},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			stockSpanner := Constructor()
			got := make([]int, len(tc.nextVals))
			for i, price := range tc.nextVals {
				got[i] = stockSpanner.Next(price)
			}
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
