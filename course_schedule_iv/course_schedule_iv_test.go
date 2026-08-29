package course_schedule_iv

import (
	"reflect"
	"testing"
)

func TestCheckIfPrerequisite(t *testing.T) {
	tt := []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		want          []bool
	}{
		{
			4,
			[][]int{{1, 0}, {2, 1}, {3, 2}},
			[][]int{{0, 1}, {3, 1}},
			[]bool{false, true},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := checkIfPrerequisite(tc.numCourses, tc.prerequisites, tc.queries)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
