package surrounded_regions

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tt := []struct {
		board [][]byte
		want  [][]byte
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X', 'O', 'O', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'X', 'O', 'O', 'X'},
				{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X'},
			},
		},
		{
			[][]byte{{'X'}},
			[][]byte{{'X'}},
		},
		{
			[][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X'},
			},
			[][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'O'},
				{'O', 'X', 'O', 'X'},
			},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			solve(tc.board)

			if !reflect.DeepEqual(tc.board, tc.want) {
				t.Fatalf("got %v want %v", tc.board, tc.want)
			}
		})
	}
}
