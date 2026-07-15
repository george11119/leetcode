package verifying_an_alien_dictionary

import "testing"

func TestIsAlienSorted(t *testing.T) {
	tt := []struct {
		words []string
		order string
		want  bool
	}{
		{
			[]string{"dag", "disk", "dog"},
			"hlabcdefgijkmnopqrstuvwxyz",
			true,
		},
		{
			[]string{"neetcode", "neet"},
			"worldabcefghijkmnpqstuvxyz",
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := isAlienSorted(tc.words, tc.order)
			want := tc.want
			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}

}
