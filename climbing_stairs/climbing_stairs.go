package climbing_stairs

//func climbStairs(n int) int {
//	if n <= 0 {
//		return 0
//	}
//
//	memo := make(map[int]int)
//
//	helper := func(i int) {
//		if i <= 1 {
//			memo[i] = 1
//			return
//		}
//
//		memo[i] = memo[i-1] + memo[i-2]
//	}
//
//	for i := range n {
//		helper(i)
//	}
//
//	return memo[n-1] + memo[n-2]
//}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	one, two := 1, 1
	for range n - 1 {
		tmp := one
		one = one + two
		two = tmp
	}

	return one
}
