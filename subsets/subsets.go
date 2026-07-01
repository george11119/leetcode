package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))

	subset := make([]int, 0, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			copySlice := make([]int, len(subset))
			copy(copySlice, subset)
			res = append(res, copySlice)
			return
		}

		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)

	return res
}
