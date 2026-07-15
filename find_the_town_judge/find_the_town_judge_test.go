package find_the_town_judge

import "testing"

func TestFindJudge(t *testing.T) {
	tt := []struct {
		n     int
		trust [][]int
		want  int
	}{
		{
			4,
			[][]int{{1, 3}, {4, 3}, {2, 3}},
			3,
		},
		{
			3,
			[][]int{{1, 3}, {2, 3}, {3, 1}, {3, 2}},
			-1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findJudge(tc.n, tc.trust)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
