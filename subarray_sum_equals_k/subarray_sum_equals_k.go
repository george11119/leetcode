package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	prefixSumCount := make(map[int]int)
	prefixSumCount[0]++

	res := 0
	prefixSum := 0

	for _, n := range nums {
		prefixSum += n
		sum := prefixSum - k

		count, ok := prefixSumCount[sum]

		if ok {
			res += count
		}

		prefixSumCount[prefixSum]++
	}

	return res
}
