package reverse_string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tt := []struct {
		s    []byte
		want []byte
	}{
		{
			[]byte{'n', 'e', 'e', 't'},
			[]byte{'t', 'e', 'e', 'n'},
		},
		{
			[]byte{'r', 'a', 'c', 'e', 'c', 'a', 'r'},
			[]byte{'r', 'a', 'c', 'e', 'c', 'a', 'r'},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := tc.s
			reverseString(got)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
