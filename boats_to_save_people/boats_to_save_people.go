package boats_to_save_people

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	i := 0
	j := len(people) - 1

	res := 0

	for i <= j {
		counter := limit
		counter -= people[j]
		j--

		if i <= j && counter-people[i] >= 0 {
			i++
		}

		res++
	}

	return res
}
