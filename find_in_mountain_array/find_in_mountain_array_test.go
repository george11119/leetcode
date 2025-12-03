package find_in_mountain_array

import "testing"

func TestFindInMountainArray(t *testing.T) {
	tt := []struct {
		mountainArr MountainArray
		target      int
		want        int
	}{
		{
			MountainArray{[]int{2, 4, 5, 2, 1}, 0},
			2,
			0,
		},
		{
			MountainArray{[]int{1, 2, 3, 4, 2, 1}, 0},
			6,
			-1,
		},
		{
			MountainArray{[]int{1, 2, 3, 4, 5, 6, 3, 2, 1}, 0},
			2,
			1,
		},
		{
			MountainArray{[]int{1, 2, 3, 5, 6, 4, 3, 2, 1}, 0},
			4,
			5,
		},
		{
			MountainArray{[]int{0, 5, 3, 1}, 0},
			1,
			3,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := findInMountainArray(tc.target, &tc.mountainArr)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
