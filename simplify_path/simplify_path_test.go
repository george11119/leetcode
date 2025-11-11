package simplify_path

import "testing"

func TestSimplifyPath(t *testing.T) {
	tt := []struct {
		path string
		want string
	}{
		{
			"/neetcode/practice//...///../courses",
			"/neetcode/practice/courses",
		},
		{
			"/..//",
			"/",
		},
		{
			"/..//_home/a/b/..///",
			"/_home/a",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := simplifyPath(tc.path)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
