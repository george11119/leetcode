package longest_repeating_character_replacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tt := []struct {
		s    string
		k    int
		want int
	}{
		{
			"XYYX", 2, 4,
		},
		{
			"AAABABB", 1, 5,
		},
		{
			"AABABBA", 1, 4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := characterReplacement(tc.s, tc.k)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
