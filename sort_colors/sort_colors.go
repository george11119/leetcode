package sort_colors

// using bucket sort
//func SortColors(nums []int) {
//	count := [3]int{}
//
//	for _, n := range nums {
//		count[n] += 1
//	}
//
//	j := 0
//	for i := 0; i < len(nums); i++ {
//		nums[i] = j
//		count[j]--
//		if count[j] == 0 {
//			j++
//		}
//	}
//}

func SortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	i := 0

	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		}
	}
}
