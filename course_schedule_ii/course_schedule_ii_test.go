package course_schedule_ii

import "testing"

func TestFindOrder(t *testing.T) {
	tt := []struct {
		numCourses    int
		prerequisites [][]int
	}{
		{
			2,
			[][]int{{1, 0}},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			findOrder(tc.numCourses, tc.prerequisites)
		})
	}
}
