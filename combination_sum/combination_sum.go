package combination_sum

func combinationSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	cur := make([]int, 0, len(nums))
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if sum == target {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		if i >= len(nums) || sum > target {
			return
		}

		cur = append(cur, nums[i])
		dfs(i, sum+nums[i])
		cur = cur[:len(cur)-1]
		dfs(i+1, sum)
	}
	dfs(0, 0)

	return res
}
