package group_anagrams

import "testing"
import "reflect"

//func TestValidAnagram(t *testing.T) {
//	tt := []struct {
//		s    string
//		t    string
//		want bool
//	}{
//		{
//			"racecar", "carrace", true,
//		},
//		{
//			"jar", "jam", false,
//		},
//	}
//
//	for _, tc := range tt {
//		t.Run("testcase", func(t *testing.T) {
//			got := validAnagram(tc.s, tc.t)
//			want := tc.want
//
//			if got != want {
//				t.Fatalf("got %v, want %v", got, want)
//			}
//		})
//	}
//}

// note: test is flaky. the results from groupAnagrams can be in different orders even when the answer is correct
func TestGroupAnagrams(t *testing.T) {
	tt := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"act", "pots", "tops", "cat", "stop", "hat"},
			[][]string{
				{"act", "cat"},
				{"pots", "tops", "stop"},
				{"hat"},
			},
		},
		{
			[]string{"x"},
			[][]string{
				{"x"},
			},
		},
		//{
		//[]string{""},
		//[][]string{
		//	{""},
		//},
		//},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			got := groupAnagrams(tc.strs)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
