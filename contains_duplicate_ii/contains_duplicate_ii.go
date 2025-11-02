package contains_duplicate_ii

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// brute force O(n^2) time
//func containsNearbyDuplicate(nums []int, k int) bool {
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < len(nums); j++ {
//			if i != j && nums[i] == nums[j] && abs(i-j) <= k {
//				return true
//			}
//		}
//	}
//	return false
//}

// hash map O(n) time O(n) space
//func containsNearbyDuplicate(nums []int, k int) bool {
//	indexOfValue := make(map[int]int)
//
//	for i := 0; i < len(nums); i++ {
//		j, ok := indexOfValue[nums[i]]
//
//		if ok {
//			if abs(j-i) <= k {
//				return true
//			}
//		}
//
//		indexOfValue[nums[i]] = i
//	}
//
//	return false
//}

// sliding window using hash set
func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool)
	i := 0

	for j := 0; j < len(nums); j++ {
		if abs(i-j) > k {
			delete(window, nums[i])
			i++
		}

		if _, ok := window[nums[j]]; ok {
			return true
		}

		window[nums[j]] = true
	}

	return false
}
