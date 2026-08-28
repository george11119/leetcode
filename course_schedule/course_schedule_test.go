package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	tt := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			2,
			[][]int{{0, 1}},
			true,
		},
		{
			2,
			[][]int{{0, 1}, {1, 0}},
			false,
		},
		{
			3,
			[][]int{{0, 1}, {2, 0}},
			true,
		},
		{
			3,
			[][]int{{0, 1}, {2, 0}, {1, 2}},
			false,
		},
		{
			6,
			[][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 5}},
			true,
		},
		{
			20,
			[][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}},
			false,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := canFinish(tc.numCourses, tc.prerequisites)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
