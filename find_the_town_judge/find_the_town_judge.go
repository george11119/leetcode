package find_the_town_judge

func findJudge(n int, trust [][]int) int {
	incoming := make(map[int]int)
	outgoing := make(map[int]int)

	for _, t := range trust {
		outgoing[t[0]] += 1
		incoming[t[1]] += 1
	}

	for p := 1; p <= n; p++ {
		if incoming[p] == n-1 && outgoing[p] == 0 {
			return p
		}
	}

	return -1
}

//func findJudge(n int, trust [][]int) int {
//	for p := 1; p <= n; p++ {
//		isJudge := true
//		trustList := make([]int, n+1)
//
//		for _, t := range trust {
//			if t[0] == p {
//				isJudge = false
//				break
//			}
//
//			if t[1] == p {
//				trustList[t[0]] = 1
//			}
//		}
//
//		if !isJudge {
//			continue
//		}
//
//		for i := 1; i <= n; i++ {
//			if i == p {
//				continue
//			}
//
//			if trustList[i] != 1 {
//				isJudge = false
//				break
//			}
//		}
//
//		if isJudge {
//			return p
//		}
//	}
//
//	return -1
//}
