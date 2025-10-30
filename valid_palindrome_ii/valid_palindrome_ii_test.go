package valid_palindrome_ii

import "testing"

func TestValidPalindrome(t *testing.T) {
	tt := []struct {
		s    string
		want bool
	}{
		{
			"aca", true,
		},
		{
			"abbadc", false,
		},
		{
			"abbda", true,
		},
		{
			"abca", true,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := validPalindrome(tc.s)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
