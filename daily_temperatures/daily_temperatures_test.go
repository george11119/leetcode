package daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tt := []struct {
		temperatures []int
		want         []int
	}{
		{
			[]int{30, 38, 30, 36, 35, 40, 28},
			[]int{1, 4, 1, 2, 1, 0, 0},
		},
		{
			[]int{22, 21, 20},
			[]int{0, 0, 0},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := dailyTemperatures(tc.temperatures)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
