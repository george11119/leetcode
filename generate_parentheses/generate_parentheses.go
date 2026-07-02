package generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	cur := make([]rune, 0, n*2)
	var dfs func(openParenCount, closeParenCount int)
	dfs = func(openParenCount, closeParenCount int) {
		if openParenCount == n {
			for i := closeParenCount; i < openParenCount; i++ {
				cur = append(cur, ')')
			}
			s := string(cur)
			res = append(res, s)
			for i := closeParenCount; i < openParenCount; i++ {
				cur = cur[:len(cur)-1]
			}
			return
		}

		cur = append(cur, '(')
		dfs(openParenCount+1, closeParenCount)
		cur = cur[:len(cur)-1]

		if openParenCount > closeParenCount {
			cur = append(cur, ')')
			dfs(openParenCount, closeParenCount+1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, 0)

	return res
}
