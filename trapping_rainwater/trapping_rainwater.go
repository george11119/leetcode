package trapping_rainwater

// waterArea := min(maxL, maxR) - height[i]

// two pointer O(n) time O(1) space
func trap(height []int) int {
	sum := 0
	l := 0
	r := len(height) - 1

	maxL := height[l]
	maxR := height[r]

	for l < r {
		if height[l] <= height[r] {
			l++
			waterArea := maxL - height[l]
			if waterArea > 0 {
				sum += waterArea
			}
			maxL = max(maxL, height[l])
		} else {
			r--
			waterArea := maxR - height[r]
			if waterArea > 0 {
				sum += waterArea
			}
			maxR = max(maxR, height[r])
		}
	}

	return sum
}

// brute force
//func trap(height []int) int {
//	sum := 0
//
//	for i := 0; i < len(height); i++ {
//		maxL := height[0]
//		maxR := height[len(height)-1]
//		for j := 0; j < i; j++ {
//			maxL = max(maxL, height[j])
//		}
//		for j := len(height) - 1; j > i; j-- {
//			maxR = max(maxR, height[j])
//		}
//		waterArea := min(maxL, maxR) - height[i]
//		if waterArea > 0 {
//			sum += waterArea
//		}
//	}
//
//	return sum
//}
