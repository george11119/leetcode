package permutations

func sliceWithIndexRemoved(arr []int, i int) []int {
	res := make([]int, 0, len(arr)-1)
	res = append(res, arr[:i]...)
	res = append(res, arr[i+1:]...)
	return res
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	cur := make([]int, 0, len(nums))
	var dfs func(vals []int)
	dfs = func(vals []int) {
		if len(cur) == len(nums) {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		for i, n := range vals {
			cur = append(cur, n)
			dfs(sliceWithIndexRemoved(vals, i))
			cur = cur[:len(cur)-1]
		}
	}
	dfs(nums)

	return res
}
