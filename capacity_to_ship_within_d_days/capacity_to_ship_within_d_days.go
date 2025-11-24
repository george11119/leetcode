package capacity_to_ship_within_d_days

// O(n logn), n = length of weights
func shipWithinDays(weights []int, days int) int {
	low := -1
	high := 0

	// low is the max value in weights, high is the sum of all weights
	for _, w := range weights {
		low = max(low, w)
		high += w
	}
	res := high + 1 // placeholder

	for low <= high {
		capacity := (low + high) / 2
		curCap := 0
		minDays := 0

		// can be optimized to stop if minDays > days
		for _, w := range weights {
			if curCap+w > capacity {
				curCap = 0
				minDays++
			}
			curCap += w
		}
		minDays++

		if minDays <= days {
			res = min(capacity, res)
			high = capacity - 1
		} else {
			low = capacity + 1
		}
	}

	return res
}
