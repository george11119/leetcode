package majority_element_ii

// O(n) time complexity, O(n) space complexity
//func majorityElement(nums []int) []int {
//	numCount := make(map[int]int)
//
//	for _, n := range nums {
//		numCount[n] += 1
//	}
//
//	res := make([]int, 0)
//	for k, v := range numCount {
//		if v > len(nums)/3 {
//			res = append(res, k)
//		}
//	}
//
//	return res
//}

// the complicated O(n) time O(1) space solution
func majorityElement(nums []int) []int {
	numCount := make(map[int]int)

	for _, n := range nums {
		numCount[n]++

		if len(numCount) >= 3 {
			for k := range numCount {
				numCount[k]--
				if numCount[k] == 0 {
					delete(numCount, k)
				}
			}
		}
	}

	res := make([]int, 0)
	for k := range numCount {
		count := 0
		for _, n := range nums {
			if n == k {
				count++
			}
		}
		if count > len(nums)/3 {
			res = append(res, k)
		}
	}

	return res
}
