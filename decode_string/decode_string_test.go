package decode_string

import "testing"

func TestDecodeString(t *testing.T) {
	tt := []struct {
		s    string
		want string
	}{
		{
			"2[a3[b]]c",
			"abbbabbbc",
		},
		{
			"2[abbb]c",
			"abbbabbbc",
		},
		{
			"axb3[z]4[c]",
			"axbzzzcccc",
		},
		{
			"ab2[c]3[d]1[x]",
			"abccdddx",
		},
		{
			"",
			"",
		},
		{
			"abc",
			"abc",
		},
		{
			"30[2[ab]]",
			"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababab",
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := decodeString(tc.s)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
