package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[l] < nums[r] {
			return nums[l]
		}

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
