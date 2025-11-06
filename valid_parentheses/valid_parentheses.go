package valid_parentheses

func checkMatchingParens(a, b rune) bool {
	if (a == '[' && b == ']') || (a == '{' && b == '}') || (a == '(' && b == ')') {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
		}
		if c == ']' || c == '}' || c == ')' {
			if len(stack) == 0 {
				return false
			}
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !checkMatchingParens(open, c) {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
