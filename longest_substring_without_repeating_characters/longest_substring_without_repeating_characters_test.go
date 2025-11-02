package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tt := []struct {
		s    string
		want int
	}{
		{
			"x",
			1,
		},
		{
			"zxyzxyz",
			3,
		},
		{
			"xxxx",
			1,
		},
		{
			"dvdf",
			3,
		},
		{
			"pwwkew",
			3,
		},
		{
			"thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsovert",
			17,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
