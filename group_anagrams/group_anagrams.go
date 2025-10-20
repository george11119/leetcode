package group_anagrams

import (
	"encoding/json"
)

func mapToString(m map[int32]int) string {
	b, _ := json.Marshal(m)
	return string(b)
}

func charCount(s string) map[int32]int {
	sCharCount := make(map[int32]int)

	for _, c := range s {
		sCharCount[c] += 1
	}

	return sCharCount
}

func groupAnagrams(strs []string) [][]string {
	anagram_slice := make([][]string, 0)
	anagram_map := make(map[string][]string)

	for _, str := range strs {
		key := mapToString(charCount(str))

		anagram_map[key] = append(anagram_map[key], str)
	}

	for _, v := range anagram_map {
		anagram_slice = append(anagram_slice, v)
	}

	return anagram_slice
}
