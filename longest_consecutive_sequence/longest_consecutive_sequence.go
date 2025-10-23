package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	inNums := make(map[int]bool)
	for _, v := range nums {
		inNums[v] = true
	}

	maxSeq := 0
	if len(nums) >= 1 {
		maxSeq++
	}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		_, ok := inNums[start-1]
		if ok {
			continue
		}
		potentialMax := 1
		inc := 1

		for {
			_, ok := inNums[start+inc]
			if !ok {
				break
			}
			potentialMax += 1
			inc += 1
			if potentialMax > maxSeq {
				maxSeq = potentialMax
			}
		}
	}

	return maxSeq
}
