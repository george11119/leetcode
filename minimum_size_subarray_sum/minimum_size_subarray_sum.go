package minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	shortest := len(nums) + 1
	sum := 0

	l := 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			shortest = min(r-l+1, shortest)
			sum -= nums[l]
			l++
		}
	}

	if shortest == len(nums)+1 {
		return 0
	}
	return shortest
}
