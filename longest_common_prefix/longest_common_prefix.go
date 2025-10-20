package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	prefix := ""

	for i, c := range strs[0] {
		for _, str := range strs {
			if len(str) <= i || string(c) != string(str[i]) {
				return prefix
			}
		}
		prefix += string(c)
	}

	return prefix
}
