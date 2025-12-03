package find_in_mountain_array

import "math"

type MountainArray struct {
	arr   []int
	count int
}

func (ma *MountainArray) get(index int) int {
	if ma.count == 100 {
		panic("oh no you made more than 100 calls!")
	}
	ma.count += 1
	return ma.arr[index]
}

func (ma *MountainArray) length() int {
	return len(ma.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	// 1. find pivot

	// should use up ~28 calls of get
	peakIdx := -1
	l := 0
	r := mountainArr.length()

	for l <= r {
		mid := (l + r) / 2
		leftOfPeak := math.MinInt
		if mid-1 > 0 {
			leftOfPeak = mountainArr.get(mid - 1)
		}
		peak := mountainArr.get(mid)
		rightOfPeak := math.MaxInt
		if mid < mountainArr.length()-1 {
			rightOfPeak = mountainArr.get(mid + 1)
		}

		if leftOfPeak < peak && peak > rightOfPeak {
			peakIdx = mid

			if peak == target {
				return peakIdx
			}

			if target > peak {
				return -1
			}

			break
		} else if leftOfPeak < peak && peak < rightOfPeak {
			l = mid + 1
		} else if leftOfPeak > peak && peak > rightOfPeak {
			r = mid - 1
		} else {
			return -1 // should never happen if mountain array is valid
		}
	}

	// 2. check left with normal binary search
	l = 0
	r = peakIdx - 1
	for l <= r {
		mid := (l + r) / 2
		valAtMid := mountainArr.get(mid)
		if valAtMid == target {
			return mid
		} else if valAtMid > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 3. check right with inverted binary search
	l = peakIdx + 1
	r = mountainArr.length() - 1

	for l <= r {
		mid := (l + r) / 2
		valAtMid := mountainArr.get(mid)
		if valAtMid == target {
			return mid
		} else if valAtMid > target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// 4. return -1 if value isnt found
	return -1
}
