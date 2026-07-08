package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	digitMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	res := make([]string, 0)
	byteDigits := []byte(digits)

	cur := make([]byte, 0, len(digits))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) && len(cur) == 0 {
			return
		}

		if i == len(digits) {
			s := string(cur)
			res = append(res, s)
			return
		}

		d := byteDigits[i]

		for _, b := range digitMap[d] {
			cur = append(cur, b)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)

	return res
}
