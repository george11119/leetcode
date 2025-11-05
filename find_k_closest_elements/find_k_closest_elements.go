package find_k_closest_elements

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// linear search + sliding window O(n) time O(1) space
//func findClosestElements(arr []int, k int, x int) []int {
//	i := 0
//	for j := 1; j < len(arr); j++ {
//		if abs(arr[j]-x) < abs(arr[i]-x) {
//			i = j
//		}
//	}
//
//	l := i
//	r := i
//
//	for r-l+1 < k {
//		if l <= 0 {
//			r++
//			continue
//		}
//		if r >= len(arr)-1 {
//			l--
//			continue
//		}
//
//		a := arr[l-1]
//		b := arr[r+1]
//
//		if abs(a-x) <= abs(b-x) {
//			l--
//		} else {
//			r++
//		}
//	}
//
//	return arr[l : r+1]
//}

// 2 pointers sliding window O(n) time O(1) space
func findClosestElements(arr []int, k int, x int) []int {
	l := 0
	r := len(arr) - 1

	for r-l+1 > k {
		a := arr[l]
		b := arr[r]

		if abs(a-x) <= abs(b-x) {
			r--
		} else {
			l++
		}
	}
	return arr[l : r+1]
}

// note: there is a more efficient way to solve this problem using binary search
