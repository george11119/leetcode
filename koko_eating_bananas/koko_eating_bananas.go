package koko_eating_bananas

// using binary search to find min speed. O(n * log m), n = number of piles, m = max value in piles
func minEatingSpeed(piles []int, h int) int {
	// used to calculate eating speed
	low := 1
	high := -1
	for _, p := range piles {
		high = max(p, high)
	}

	// minimum speed
	res := high + 1

	for low <= high {
		minH := 0
		speed := (low + high) / 2

		for _, pile := range piles {
			minH += pile / speed
			if pile%speed != 0 {
				minH++
			}
		}

		// binary search to find next speed
		if minH <= h {
			res = min(speed, res)
			high = speed - 1
		} else {
			low = speed + 1
		}
	}

	return res
}

// brute force O(n * m), n = number of piles, m = max value in piles
//func minEatingSpeed(piles []int, h int) int {
//	speed := 1
//	i := 0
//	minH := 0
//
//	for {
//		eaten := 0
//
//    // am i stupid why did I implement division with a for loop
//		for i < len(piles) {
//			eaten += speed
//			if eaten >= piles[i] {
//				i++
//				eaten = 0
//			}
//			minH++
//		}
//
//		if minH <= h {
//			return speed
//		} else {
//			speed++
//			minH = 0
//			i = 0
//		}
//	}
//}
