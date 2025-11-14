package largest_rectangle_in_histogram

type Pair struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	s := make([]Pair, 0, len(heights))

	for i, h := range heights {
		begin := i
		for len(s) > 0 && s[len(s)-1].height > h {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			m := top.height * (i - top.index)
			maxArea = max(m, maxArea)
			begin = top.index
		}
		s = append(s, Pair{begin, h})
	}

	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		m := top.height * (len(heights) - top.index)
		maxArea = max(m, maxArea)
	}

	return maxArea
}
