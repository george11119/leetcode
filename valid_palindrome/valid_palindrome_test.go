package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		s    string
		want bool
	}{
		{
			"Was it a car or a cat I saw?",
			true,
		},
		{
			"tab a cat",
			false,
		},
		{
			"?ab?????",
			false,
		},
		{
			".,",
			true,
		},
		{
			"0P",
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := isPalindrome(tc.s)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
