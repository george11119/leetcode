package first_missing_positive

// O(n) time O(n) space
//func firstMissingPositive(nums []int) int {
//	smallest := 1
//
//	numMap := make(map[int]bool)
//
//	for _, n := range nums {
//		numMap[n] = true
//	}
//
//	for {
//		_, ok := numMap[smallest]
//		if !ok {
//			break
//		}
//		smallest++
//	}
//
//	return smallest
//}

func absolute(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// O(n) time, O(1) space assuming nums array is given
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}

	for _, n := range nums {
		i := absolute(n) - 1

		if i < 0 || i >= len(nums) {
			continue
		}

		if nums[i] > 0 {
			nums[i] = nums[i] * -1
		} else if nums[i] == 0 {
			nums[i] = -(len(nums) + 1)
		}
	}

	smallest := 1
	for smallest <= len(nums) {
		i := smallest - 1
		if nums[i] >= 0 {
			break
		}
		smallest++
	}

	return smallest
}
