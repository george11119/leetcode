package container_with_most_water

func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1
	resMax := 0

	for i < j {
		area := min(heights[i], heights[j]) * (j - i)
		if area > resMax {
			resMax = area
		}

		if heights[i] > heights[j] {
			j--
		} else {
			i++
		}
	}

	return resMax
}
