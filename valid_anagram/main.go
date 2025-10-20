package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	s_seen := make(map[int32]int)
	t_seen := make(map[int32]int)

	for _, c := range s {
		s_seen[c] += 1
	}
	for _, c := range t {
		t_seen[c] += 1
	}

	return reflect.DeepEqual(s_seen, t_seen)
}

func main() {
	fmt.Println(isAnagram("racecar", "carrace"))
	fmt.Println(isAnagram("jar", "jam"))
}
