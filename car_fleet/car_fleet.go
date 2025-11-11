package car_fleet

import (
	"sort"
)

type car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]car, 0, n)
	s := make([]car, 0, n)

	for i := 0; i < n; i++ {
		cars = append(cars, car{position[i], speed[i]})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	for _, c := range cars {
		if len(s) == 0 {
			s = append(s, c)
			continue
		}

		prev := s[len(s)-1]
		s = append(s, c)

		timePrev := float32(target-prev.position) / float32(prev.speed)
		timeC := float32(target-c.position) / float32(c.speed)

		if timeC <= timePrev {
			s = s[:len(s)-1]
		}
	}

	return len(s)
}
