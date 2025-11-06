package baseball_game

import "testing"

func TestCalPoints(t *testing.T) {
	tt := []struct {
		operations []string
		want       int
	}{
		{
			[]string{"1", "2", "+", "C", "5", "D"},
			18,
		},
		{
			[]string{"5", "D", "+", "C"},
			15,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := calPoints(tc.operations)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
