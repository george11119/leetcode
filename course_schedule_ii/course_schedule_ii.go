package course_schedule_ii

import "slices"

func findOrder(numCourses int, prerequisites [][]int) []int {
	prereqGraph := make(map[int][]int)
	for _, pr := range prerequisites {
		a := pr[0]
		b := pr[1]

		prereqGraph[b] = append(prereqGraph[b], a)
	}

	res := make([]int, 0)

	visiting := make(map[int]bool)
	visited := make(map[int]bool)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if visiting[i] {
			return false
		}

		if visited[i] {
			return true
		}

		visiting[i] = true
		visited[i] = true

		for _, j := range prereqGraph[i] {
			if !dfs(j) {
				return false
			}
		}

		visiting[i] = false
		res = append(res, i)

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return []int{}
		}
	}

	slices.Reverse(res)
	return res
}
