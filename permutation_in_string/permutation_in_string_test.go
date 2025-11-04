package permutation_in_string

import "testing"

func TestCheckInclusion(t *testing.T) {
	tt := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			"abc",
			"baxyzabc",
			true,
		},
		{
			"abc",
			"lecabee",
			true,
		},
		{
			"abc",
			"lecaabe",
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := checkInclusion(tc.s1, tc.s2)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
