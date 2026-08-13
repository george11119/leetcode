package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tt := []struct {
		n    int
		want int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
		{
			44,
			1134903170,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := climbStairs(tc.n)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
