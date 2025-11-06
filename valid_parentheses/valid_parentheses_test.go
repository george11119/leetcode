package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tt := []struct {
		s    string
		want bool
	}{
		{
			"[]",
			true,
		},
		{
			"([{}])",
			true,
		},
		{
			"[(])",
			false,
		},
		{
			"(",
			false,
		},
		{
			")",
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := isValid(tc.s)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
