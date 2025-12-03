package median_of_two_sorted_arrays

import (
	"math"
)

// binary search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var smallerNums []int
	var largerNums []int
	if len(nums1) > len(nums2) {
		smallerNums = nums2
		largerNums = nums1
	} else {
		smallerNums = nums1
		largerNums = nums2
	}

	l := 0
	r := len(smallerNums)

	for l <= r {
		idx1 := (l + r) / 2
		idx2 := ((len(smallerNums) + len(largerNums) + 1) / 2) - idx1

		minIntSmaller := math.MinInt
		if idx1-1 >= 0 {
			minIntSmaller = smallerNums[idx1-1]
		}

		maxIntSmaller := math.MaxInt
		if idx1 <= len(smallerNums)-1 {
			maxIntSmaller = smallerNums[idx1]
		}

		minIntLarger := math.MinInt
		if idx2-1 >= 0 {
			minIntLarger = largerNums[idx2-1]
		}

		maxIntLarger := math.MaxInt
		if idx2 <= len(largerNums)-1 {
			maxIntLarger = largerNums[idx2]
		}

		if minIntSmaller <= maxIntLarger && minIntLarger <= maxIntSmaller {
			if (len(nums1)+len(nums2))%2 == 0 {
				return float64(max(minIntSmaller, minIntLarger)+min(maxIntSmaller, maxIntLarger)) / 2
			} else {
				return float64(max(minIntSmaller, minIntLarger))
			}
		} else if maxIntSmaller < minIntLarger {
			l = idx1 + 1
		} else {
			r = idx1 - 1
		}
	}

	return -1
}

// brute force
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums3 := make([]int, 0, len(nums1)+len(nums2))
//
//	i := 0
//	j := 0
//
//	for i < len(nums1) && j < len(nums2) {
//		if nums1[i] < nums2[j] {
//			nums3 = append(nums3, nums1[i])
//			i++
//		} else {
//			nums3 = append(nums3, nums2[j])
//			j++
//		}
//	}
//
//	for i < len(nums1) {
//		nums3 = append(nums3, nums1[i])
//		i++
//	}
//
//	for j < len(nums2) {
//		nums3 = append(nums3, nums2[j])
//		j++
//	}
//
//	if len(nums3)%2 == 0 {
//		median := float64(nums3[len(nums3)/2-1]+nums3[len(nums3)/2]) / 2
//		return median
//	} else {
//		median := float64(nums3[len(nums3)/2])
//		return median
//	}
//}
