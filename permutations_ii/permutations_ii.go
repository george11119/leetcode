package permutations_ii

//func sliceWithIndexRemoved(arr []int, i int) []int {
//	res := make([]int, 0, len(arr)-1)
//	res = append(res, arr[:i]...)
//	res = append(res, arr[i+1:]...)
//	return res
//}
//
//func permuteUnique(nums []int) [][]int {
//	res := make([][]int, 0)
//	sort.Ints(nums)
//
//	cur := make([]int, 0, len(nums))
//	var dfs func(vals []int)
//	dfs = func(vals []int) {
//		if len(cur) == len(nums) {
//			copySlice := make([]int, len(cur))
//			copy(copySlice, cur)
//			res = append(res, copySlice)
//			return
//		}
//
//		for i := 0; i < len(vals); i++ {
//			cur = append(cur, vals[i])
//			for i+1 < len(vals) && vals[i] == vals[i+1] {
//				i++
//			}
//			dfs(sliceWithIndexRemoved(vals, i))
//			cur = cur[:len(cur)-1]
//		}
//	}
//	dfs(nums)
//
//	return res
//}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	count := make(map[int]int)

	for _, n := range nums {
		count[n]++
	}

	cur := make([]int, 0, len(nums))
	var dfs func()
	dfs = func() {
		if len(cur) == len(nums) {
			copySlice := make([]int, len(cur))
			copy(copySlice, cur)
			res = append(res, copySlice)
			return
		}

		for n, _ := range count {
			if count[n] > 0 {
				cur = append(cur, n)
				count[n]--
				dfs()
				count[n]++
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs()

	return res
}
