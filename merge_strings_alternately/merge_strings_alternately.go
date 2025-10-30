package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	res := ""

	for i < len(word1) && j < len(word2) {
		res += string(word1[i])
		res += string(word2[j])
		i++
		j++
	}

	for l := i; l < len(word1); l++ {
		res += string(word1[l])
	}

	for l := j; l < len(word2); l++ {
		res += string(word2[l])
	}

	return res
}
