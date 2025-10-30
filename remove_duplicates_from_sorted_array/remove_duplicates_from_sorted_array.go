package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	k := 1

	i := 0
	j := 0

	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			k++
			i++
			nums[i] = nums[j]
		}
	}

	return k
}
