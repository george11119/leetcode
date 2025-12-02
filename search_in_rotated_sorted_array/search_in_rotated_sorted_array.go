package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[l] < nums[r] {
			break
		}

		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	// l is now the index that points at the lowest value in nums
	r = len(nums) - 1
	if !(nums[l] <= target && target <= nums[r]) {
		r = l - 1
		l = 0
	}

	res := -1
	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
