package valid_palindrome

func isAlphabet(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j && !isAlphabet(s[i]) {
			i++
		}

		for i < j && !isAlphabet(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}
