package remove_element

func removeElement(nums *[]int, val int) int {
	n := *nums
	k := 0

	for i := 0; i < len(n); i++ {
		if n[i] != val {
			n[k] = n[i]
			k++
		}
	}

	*nums = n
	return k
}
