package encoding_and_decoding

import (
	"reflect"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	tt := []struct {
		input []string
		want  []string
	}{
		{
			[]string{"hello", "world"},
			[]string{"hello", "world"},
		},
		{
			[]string{"4#8324", "world"},
			[]string{"4#8324", "world"},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			encodedString := Encode(tc.input)
			decodedString := Decode(encodedString)

			if !reflect.DeepEqual(decodedString, tc.want) {
				t.Fatalf("got %v want %v", decodedString, tc.want)
			}
		})
	}
}
