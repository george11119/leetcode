package longest_repeating_character_replacement

// O(26 * n) time O(1) space
func characterReplacement(s string, k int) int {
	freq := [26]int{}
	l := 0
	res := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]-'A']++
		maxFreq := 0
		for _, i := range freq {
			maxFreq = max(maxFreq, i)
		}

		for (r-l+1)-maxFreq > k {
			freq[s[l]-'A']--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}

// optimal answer O(n) time O(1) space
//func characterReplacement(s string, k int) int {
//	freq := [26]int{}
//	l := 0
//	res := 0
//	maxFreq := 0
//
//	for r := 0; r < len(s); r++ {
//		freq[s[r]-'A']++
//		maxFreq = max(freq[s[r]-'A'], maxFreq)
//
//		for (r-l+1)-maxFreq > k {
//			freq[s[l]-'A']--
//			l++
//		}
//
//		res = max(res, r-l+1)
//	}
//
//	return res
//}
