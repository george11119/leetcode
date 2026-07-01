package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	cur := make([]int, 0, k)
	var dfs func(i int)
	dfs = func(i int) {
		if len(cur) == k {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		if i > n {
			return
		}

		if len(cur) < k {
			cur = append(cur, i)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
		dfs(i + 1)
	}
	dfs(1)

	return res
}
