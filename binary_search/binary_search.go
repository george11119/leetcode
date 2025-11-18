package binary_search

// binary search
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			return -1
		}

		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
}

// linear search
//func search(nums []int, target int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == target {
//			return i
//		}
//	}
//
//	return -1
//}
