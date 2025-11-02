package longest_substring_without_repeating_characters

// O(n) time O(n) space
//
//	func lengthOfLongestSubstring(s string) int {
//		window := make(map[rune]bool)
//		l := 0
//		res := 0
//
//		for r := 0; r < len(s); r++ {
//			c := rune(s[r])
//
//			for {
//				// there is no duplicate in the window's substring
//				if _, ok := window[c]; !ok {
//					break
//				}
//				delete(window, rune(s[l]))
//				l++
//			}
//
//			window[c] = true
//
//			res = max(res, r-l+1)
//		}
//
//		return res
//	}

// O(n) time O(m) space, where m is the length of the longest substring
func lengthOfLongestSubstring(s string) int {
	charCount := make(map[uint8]int)
	l := 0
	res := 0

	for r := 0; r < len(s); r++ {
		if _, ok := charCount[s[r]]; ok {
			l = max(l, charCount[s[r]]+1)
		}

		charCount[s[r]] = r

		res = max(res, r-l+1)
	}

	return res
}
