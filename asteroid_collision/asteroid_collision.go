package asteroid_collision

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func asteroidCollision(asteroids []int) []int {
	s := make([]int, 0)

	for _, a := range asteroids {
		for len(s) > 0 && s[len(s)-1] > 0 && a < 0 {
			diff := abs(s[len(s)-1]) - abs(a)

			if diff < 0 {
				s = s[:len(s)-1]
			} else if diff > 0 {
				a = 0
			} else {
				a = 0
				s = s[:len(s)-1]
			}
		}

		if a != 0 {
			s = append(s, a)
		}
	}

	return s
}
