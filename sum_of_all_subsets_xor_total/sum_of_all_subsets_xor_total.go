package sum_of_all_subsets_xor_total

func calculateXorSum(subset []int) int {
	res := 0
	for _, n := range subset {
		res ^= n
	}

	return res
}

func subsetXORSum(nums []int) int {
	res := 0

	subset := make([]int, 0, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			res += calculateXorSum(subset)
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
