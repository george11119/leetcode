package sliding_window_maximum

import "fmt"

func debugPrint(nums []int, indexArr []int) {
	debug := make([]int, 0, len(indexArr))
	for _, i := range indexArr {
		debug = append(debug, nums[i])
	}
	fmt.Println("index: ", indexArr)
	fmt.Println("value: ", debug)
}

// efficient O(n) time O(n) space solution using deques
func maxSlidingWindow(nums []int, k int) []int {
	l := 0
	r := 0
	res := make([]int, 0, len(nums)-k+1)
	q := make([]int, 0, len(nums))

	// loop until r hits end
	for r < len(nums) {
		//fmt.Println(l)

		// make sure first index in deque is still in range of the window of l and r
		for len(q) > 0 && l > q[0] {
			q = q[1:]
		}

		// add num indexes into deque. Adds each index exactly once
		for (r - l + 1) <= k {
			// remove indexes at end of deque which correspond to values smaller than
			// value we are trying to insert
			for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
				q = q[:len(q)-1]
			}

			// add index to end
			q = append(q, r)
			r++
		}

		// add to result
		res = append(res, nums[q[0]])
		//debugPrint(nums, q)
		l++
	}

	return res
}

// O(k * n) inefficient solution
//func maxSlidingWindow(nums []int, k int) []int {
//	l := 0
//	r := k - 1
//	res := make([]int, len(nums)-k+1)
//
//	maxVal := nums[l]
//	maxValIdx := -1
//
//	for r < len(nums) {
//		if l < maxValIdx && nums[r] >= maxVal {
//			maxVal = nums[r]
//			maxValIdx = r
//		} else {
//			maxVal = nums[l]
//			maxValIdx = l
//			for i := l; i <= r; i++ {
//				if nums[i] >= maxVal {
//					maxVal = nums[i]
//					maxValIdx = i
//				}
//			}
//		}
//		res[l] = maxVal
//		l++
//		r++
//	}
//
//	return res
//}
