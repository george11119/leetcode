package minimum_window_substring

// triple for loop 😎
//func minWindow(s string, t string) string {
//	minLen := len(s) + 1
//	res := ""
//
//	for i := 0; i < len(s); i++ {
//		for j := i; j < len(s); j++ {
//			subStr := s[i : j+1]
//
//			tChars := make(map[rune]int)
//			for _, d := range t {
//				tChars[d]++
//			}
//			for _, c := range subStr {
//				if _, ok := tChars[c]; ok {
//					tChars[c]--
//					if tChars[c] == 0 {
//						delete(tChars, c)
//					}
//				}
//
//				if len(tChars) == 0 && len(subStr) < minLen {
//					minLen = len(subStr)
//					res = subStr
//				}
//			}
//		}
//	}
//
//	return res
//}

// optimal sliding window
func minWindow(s string, t string) string {
	res := ""
	minLen := len(s) + 1

	tChars := make(map[rune]int)
	window := make(map[rune]int)
	for _, c := range t {
		window[c] = 0
		tChars[c]++
	}
	have := 0
	need := len(tChars)

	l := 0
	for r, c := range s {
		if _, ok := window[c]; ok {
			window[c]++

			if window[c] == tChars[c] {
				have++
			}
		}

		for have == need {
			d := rune(s[l])
			if _, ok := window[d]; ok {
				window[d]--
				if window[d] < tChars[d] {
					have--
					if minLen > r-l+1 {
						minLen = r - l + 1
						res = s[l : r+1]
					}
				}
			}
			l++
		}
	}

	return res
}
