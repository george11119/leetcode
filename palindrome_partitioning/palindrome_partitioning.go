package palindrome_partitioning

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func partition(s string) [][]string {
	res := make([][]string, 0)

	cur := make([]int, 0, len(s))
	var dfs func(i int)
	dfs = func(i int) {
		// split string and check for palindrome partition
		if i >= len(s) {
			// split string into all possible partitions
			strSplit := make([]string, 0, len(cur)+1)
			prev := 0
			for _, j := range cur {
				strSplit = append(strSplit, s[prev:j])
				prev = j
			}
			strSplit = append(strSplit, s[prev:])

			// check for palindrome
			allPalindrome := true
			for _, split := range strSplit {
				if !isPalindrome(split) {
					allPalindrome = false
					break
				}
			}

			if allPalindrome {
				res = append(res, strSplit)
			}

			return
		}

		cur = append(cur, i)
		dfs(i + 1)
		cur = cur[:len(cur)-1]
		dfs(i + 1)
	}
	dfs(1)

	return res
}

//func isPalindrome(s string, i, j int) bool {
//	for i < j {
//		if s[i] != s[j] {
//			return false
//		}
//		i++
//		j--
//	}
//
//	return true
//}
//
//func partition(s string) [][]string {
//	res := make([][]string, 0)
//	part := make([]string, 0, len(s))
//
//	var dfs func(i int)
//	dfs = func(i int) {
//		if i >= len(s) {
//			res = append(res, append([]string{}, part...))
//			return
//		}
//
//		for j := i; j < len(s); j++ {
//			if isPalindrome(s, i, j) {
//				part = append(part, s[i:j+1])
//				dfs(j + 1)
//				part = part[:len(part)-1]
//			}
//		}
//	}
//	dfs(0)
//
//	return res
//}
