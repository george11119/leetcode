package guess_number_higher_or_lower

// note: remove the pick from this function when submitting into leetcode
func guess(n, pick int) int {
	if n > pick {
		return -1
	} else if n < pick {
		return 1
	} else {
		return 0
	}
}

// note: remove the pick from this function when submitting into leetcode
func guessNumber(n, pick int) int {
	left := 1
	right := n

	for {
		mid := (left + right) / 2
		guessRes := guess(mid, pick)

		if guessRes == -1 {
			right = mid - 1
		} else if guessRes == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
}
