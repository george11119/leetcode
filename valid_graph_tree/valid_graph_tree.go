package valid_graph_tree

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}

	graphMap := make(map[int][]int)
	for _, e := range edges {
		u := e[0]
		v := e[1]

		graphMap[u] = append(graphMap[u], v)
		graphMap[v] = append(graphMap[v], u)
	}

	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	var dfs func(i, prev int) bool
	dfs = func(i, prev int) bool {
		if visiting[i] {
			return false
		}

		if visited[i] {
			return true
		}

		visited[i] = true
		visiting[i] = true

		for _, j := range graphMap[i] {
			if prev != j && !dfs(j, i) {
				return false
			}
		}

		visiting[i] = false

		return true
	}

	if !dfs(0, 0) {
		return false
	}

	for range visited {
		n--
	}

	return n == 0
}
