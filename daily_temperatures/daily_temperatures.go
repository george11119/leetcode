package daily_temperatures

type Pair struct {
	temp  int
	index int
}

// using stack ordered in decreasing order O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := make([]Pair, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(s) > 0 && s[len(s)-1].temp < temperatures[i] {
			prev := s[len(s)-1]
			s = s[:len(s)-1]
			res[prev.index] = i - prev.index
		}
		s = append(s, Pair{temperatures[i], i})
	}

	return res
}

// DP O(n)
//func dailyTemperatures(temperatures []int) []int {
//	res := make([]int, len(temperatures))
//	n := len(temperatures)
//
//	for i := n - 2; i >= 0; i-- {
//		j := i + 1
//
//		for j < n && temperatures[i] >= temperatures[j] {
//			if res[j] == 0 {
//				j = n
//				break
//			}
//			j += res[j]
//		}
//
//		if j < n {
//			res[i] = j - i
//		}
//	}
//
//	return res
//}

// brute force O(n^2)
//func dailyTemperatures(temperatures []int) []int {
//	res := make([]int, len(temperatures))
//
//	for i := 0; i < len(temperatures); i++ {
//		for j := i; j < len(temperatures); j++ {
//			if temperatures[j] > temperatures[i] {
//				res[i] = j - i
//				break
//			}
//			res[i] = 0
//		}
//	}
//
//	return res
//}
