package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	tt := []struct {
		s    string
		t    string
		want string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"OUZODYXAZV",
			"XYZ",
			"YXAZ",
		},
		{
			"xyz",
			"xyz",
			"xyz",
		},
		{
			"x",
			"xy",
			"",
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := minWindow(tc.s, tc.t)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
