package search_insert_position

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for {
		if left > right {
			if nums[mid] > target {
				return max(0, mid)
			} else {
				return mid + 1
			}
		}

		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
}
