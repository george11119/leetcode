package valid_palindrome_ii

// valid palindrome check
func helper(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return helper(s, l+1, r) || helper(s, l, r-1)
		}
		l++
		r--
	}

	return true
}
