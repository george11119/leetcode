package sqrt_x

import "testing"

func TestMySqrt(t *testing.T) {
	tt := []struct {
		x    int
		want int
	}{
		{
			9, 3,
		},
		{
			13, 3,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := mySqrt(tc.x)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
