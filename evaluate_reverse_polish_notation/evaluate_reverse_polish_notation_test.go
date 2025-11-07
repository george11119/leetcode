package evaluate_reverse_polish_notation

import "testing"

func TestEvalRPN(t *testing.T) {
	tt := []struct {
		tokens []string
		want   int
	}{
		{
			[]string{"1", "2", "+", "3", "*", "4", "-"},
			5,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := evalRPN(tc.tokens)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
