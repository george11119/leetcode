package partition_to_k_equal_subsets

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%k != 0 {
		return false
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	used := make([]bool, len(nums))
	subsetSum := sum / k

	var dfs func(i, k, cur int) bool
	dfs = func(i, k, cur int) bool {
		if k == 0 {
			return true
		}

		if cur == subsetSum {
			return dfs(0, k-1, 0)
		}

		for j := i; j < len(nums); j++ {
			if used[j] || cur+nums[j] > subsetSum {
				continue
			}

			used[j] = true
			if dfs(j+1, k, cur+nums[j]) {
				return true
			}
			used[j] = false

			if cur == 0 {
				return false
			}
		}

		return false
	}

	return dfs(0, k, 0)
}
