package matchsticks_to_square

import (
	"sort"
)

func makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, v := range matchsticks {
		totalLen += v
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	buckets := []int{0, 0, 0, 0}
	sideLen := totalLen / 4

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i >= len(matchsticks) {
			return true
		}

		for j := 0; j < len(buckets); j++ {
			if buckets[j]+matchsticks[i] <= sideLen {
				buckets[j] += matchsticks[i]
				if dfs(i + 1) {
					return true
				}
				buckets[j] -= matchsticks[i]
			}

			if buckets[j] == 0 {
				return false
			}
		}

		return false
	}

	return dfs(0)
}
