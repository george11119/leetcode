package boats_to_save_people

import "testing"

func TestNumRescueBoats(t *testing.T) {
	tt := []struct {
		people []int
		limit  int
		want   int
	}{
		{
			[]int{5, 1, 4, 2},
			6,
			2,
		},
		{
			[]int{1, 3, 2, 3, 2},
			3,
			4,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := numRescueBoats(tc.people, tc.limit)
			want := tc.want

			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
