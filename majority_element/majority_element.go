package majority_element

// hash map solution
//func majorityElement(nums []int) int {
//	count := make(map[int]int)
//	res := 0
//	maxCount := 0
//
//	for _, n := range nums {
//		count[n] += 1
//
//		if count[n] > maxCount {
//			res = n
//			maxCount = count[n]
//		}
//
//	}
//
//	return res
//}

// boyer moore algorithm solution
func majorityElement(nums []int) int {
	res := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			res = num
		}

		if num == res {
			count++
		} else {
			count--
		}
	}

	return res
}
