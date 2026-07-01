package combination_sum_ii

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)

	cur := make([]int, 0, len(candidates))
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if sum == target {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		if i >= len(candidates) || sum > target {
			return
		}

		cur = append(cur, candidates[i])
		dfs(i+1, sum+candidates[i])
		cur = cur[:len(cur)-1]

		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}

		dfs(i+1, sum)
	}
	dfs(0, 0)

	return res
}
