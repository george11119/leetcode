package permutation_in_string

// O(n)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	matches := 0
	s1CharCount := [26]int{}
	s2CharCount := [26]int{}

	for i := range s1 {
		s1CharCount[s1[i]-'a']++
		s2CharCount[s2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if s1CharCount[i] == s2CharCount[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		idx := s2[l] - 'a'
		s2CharCount[idx]--
		// if the decremented value previously matched, remove it from matches
		if s2CharCount[idx]+1 == s1CharCount[idx] {
			matches--
		}
		// if the decremented value now matches, add it to matches
		if s2CharCount[idx] == s1CharCount[idx] {
			matches++
		}
		l++

		idx = s2[r] - 'a'
		s2CharCount[idx]++
		// if the incremented value previously matched, remove it from matches
		if s2CharCount[idx]-1 == s1CharCount[idx] {
			matches--
		}
		// if the incremented value now matches, add it to matches
		if s2CharCount[idx] == s1CharCount[idx] {
			matches++
		}
	}

	return matches == 26
}
