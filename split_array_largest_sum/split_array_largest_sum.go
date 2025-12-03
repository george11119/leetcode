package split_array_largest_sum

func getMax(nums []int) int {
	res := nums[0]
	for _, n := range nums {
		res = max(n, res)
	}

	return res
}

func getSum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}

	return res
}

func canSplit(nums []int, k int, minVal int) bool {
	sum := 0
	subArrCount := 1

	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > minVal {
			sum = 0
			subArrCount++
		}
		sum += nums[i]
		if subArrCount > k {
			return false
		}
	}

	return true
}

func splitArray(nums []int, k int) int {
	low := getMax(nums)
	high := getSum(nums)

	for low <= high {
		mid := (low + high) / 2
	
		if canSplit(nums, k, mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
