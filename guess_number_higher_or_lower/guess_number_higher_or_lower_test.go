package guess_number_higher_or_lower

import "testing"

func TestGuessNumber(t *testing.T) {
	tt := []struct {
		n    int
		pick int
		want int
	}{
		{
			5, 3, 3,
		},
		{
			15, 10, 10,
		},
		{
			1, 1, 1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := guessNumber(tc.n, tc.pick)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
