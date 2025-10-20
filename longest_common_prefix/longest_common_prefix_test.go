package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tt := []struct {
		name string
		strs []string
		want string
	}{
		{
			"case 1", []string{"bat", "bag", "bank", "band"}, "ba",
		},
		{
			"case 2", []string{"dance", "dag", "danger", "damage"}, "da",
		},
		{
			"case 3", []string{"neet", "feet"}, "",
		},
		{
			"case 4", []string{"abc", "", "abcd"}, "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := longestCommonPrefix(tc.strs)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
