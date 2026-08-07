package open_the_lock

import "testing"

func TestOpenLock(t *testing.T) {
	tt := []struct {
		deadends []string
		target   string
		want     int
	}{
		{
			[]string{"1111", "0120", "2020", "3333"},
			"5555",
			20,
		},
		{
			[]string{"4443", "4445", "4434", "4454", "4344", "4544", "3444", "5444"},
			"4444",
			-1,
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := openLock(tc.deadends, tc.target)
			want := tc.want

			if got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
