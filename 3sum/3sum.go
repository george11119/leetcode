package _3sum

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
