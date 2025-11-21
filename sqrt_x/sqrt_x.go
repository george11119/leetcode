package sqrt_x

func mySqrt(x int) int {
	left := 0
	right := x
	res := 0

	for left <= right {
		mid := (left + right) / 2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
			res = mid
		} else {
			return mid
		}
	}

	return res
}

// brute force
//func mySqrt(x int) int {
//	for i := 1; i < x; i++ {
//		res := i * i
//		if res > x {
//			return i - 1
//		}
//	}
//
//	return -1
//}
