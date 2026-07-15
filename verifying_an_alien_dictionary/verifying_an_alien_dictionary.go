package verifying_an_alien_dictionary

// returns true if word1 comes before word2
func checkLexographicOrder(word1, word2 string, alienMap map[rune]int) bool {
	m := len(word1)
	n := len(word2)

	for i := 0; i < m && i < n; i++ {
		if word1[i] == word2[i] {
			continue
		}

		return alienMap[rune(word1[i])] < alienMap[rune(word2[i])]
	}

	return m < n
}

func isAlienSorted(words []string, order string) bool {
	while := make(map[rune]int)
	for i, letter := range order {
		while[letter] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !checkLexographicOrder(words[i], words[i+1], while) {
			return false
		}
	}

	return true
}
