package merge_strings_alternately

import "testing"

func TestMergeAlternately(t *testing.T) {
	tt := []struct {
		word1 string
		word2 string
		want  string
	}{
		{
			"abc",
			"xyz",
			"axbycz",
		},
		{
			"ab",
			"abbxxc",
			"aabbbxxc",
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := mergeAlternately(tc.word1, tc.word2)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
