package word_search

import (
	"testing"
)

func TestWordSearch(t *testing.T) {
	tt := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'D'},
				{'S', 'A', 'A', 'T'},
				{'A', 'C', 'A', 'E'},
			},
			"CAT",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				{'A'},
			},
			"A",
			true,
		},
		{
			[][]byte{
				{'A', 'B'},
			},
			"ABAB",
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := exist(tc.board, tc.word)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
