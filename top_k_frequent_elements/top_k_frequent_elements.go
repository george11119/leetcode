package top_k_frequent_elements

// bucket sort where the count array index is number of times a value appears
// and values in the array are arrays that contain the values
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, v := range nums {
		count[v] += 1
	}

	freq := make([][]int, len(nums)+1)

	for v, c := range count {
		freq[c] = append(freq[c], v)
	}

	res := make([]int, 0, k)

	for i := len(freq) - 1; i > 0; i-- {
		if len(freq[i]) != 0 {
			for _, v := range freq[i] {
				res = append(res, v)
				if len(res) == k {
					return res
				}
			}
		}
	}

	return res
}
