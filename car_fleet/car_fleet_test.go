package car_fleet

import "testing"

func TestCarFleet(t *testing.T) {
	tt := []struct {
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			10,
			[]int{7, 5, 3},
			[]int{1, 2, 3},
			1,
		},
		{
			10,
			[]int{1, 4},
			[]int{3, 2},
			1,
		},
		{
			10,
			[]int{4, 1, 0, 7},
			[]int{2, 2, 1, 1},
			3,
		},
		{
			31,
			[]int{5, 26, 18, 25, 29, 21, 22, 12, 19, 6},
			[]int{7, 6, 6, 4, 3, 4, 9, 7, 6, 4},
			6,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := carFleet(tc.target, tc.position, tc.speed)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
