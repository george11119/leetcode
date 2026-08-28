package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	prereqGraph := make(map[int][]int)
	for _, pr := range prerequisites {
		a := pr[0]
		b := pr[1]

		prereqGraph[a] = append(prereqGraph[a], b)
	}

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
			if dfs(j) == false {
				return false
			}
		}

		visiting[i] = false

		return true
	}

	for i := range numCourses {
		if !visited[i] && !dfs(i) {
			return false
		}
	}

	return true
}
