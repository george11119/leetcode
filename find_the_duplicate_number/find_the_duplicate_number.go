package find_the_duplicate_number

// O(n) time, O(n) space
//func findDuplicate(nums []int) int {
//	duplicateMap := make(map[int]int)
//
//	for _, n := range nums {
//		if duplicateMap[n] != 0 {
//			return n
//		}
//
//		duplicateMap[n] += 1
//	}
//
//	return -1
//}

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]

		if slow == slow2 {
			return slow
		}
	}
}
