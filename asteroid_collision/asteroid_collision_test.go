package asteroid_collision

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tt := []struct {
		asteroids []int
		want      []int
	}{
		{
			[]int{2, 4, -4, -1},
			[]int{2},
		},
		{
			[]int{5, 5},
			[]int{5, 5},
		},
		{
			[]int{7, -3, 9},
			[]int{7, 9},
		},
		{
			[]int{10, 2, -5},
			[]int{10},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := asteroidCollision(tc.asteroids)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
