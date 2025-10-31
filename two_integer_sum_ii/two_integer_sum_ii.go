package two_integer_sum_ii

// brute force
//func twoSum(numbers []int, target int) []int {
//	for i := 0; i < len(numbers); i++ {
//		for j := 0; j < len(numbers); j++ {
//			if i >= j {
//				continue
//			}
//			sum := numbers[i] + numbers[j]
//			if sum == target {
//				return []int{i + 1, j + 1}
//			}
//		}
//	}
//
//	// should never reach here
//	return []int{}
//}

// two pointer O(n) time O(1) space
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
