package subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	cur := make([]int, 0, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		cur = append(cur, nums[i])
		dfs(i + 1)
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		cur = cur[:len(cur)-1]
		dfs(i + 1)
	}
	dfs(0)

	return res
}
