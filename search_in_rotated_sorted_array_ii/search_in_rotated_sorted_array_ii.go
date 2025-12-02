package search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[mid+1] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			r--
		}
	}

	return target == nums[l]
}
