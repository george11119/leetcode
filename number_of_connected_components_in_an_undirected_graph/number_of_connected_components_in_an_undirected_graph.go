package number_of_connected_components_in_an_undirected_graph

func countComponents(n int, edges [][]int) int {
	graphMap := make(map[int][]int)

	for _, e := range edges {
		u := e[0]
		v := e[1]

		graphMap[u] = append(graphMap[u], v)
		graphMap[v] = append(graphMap[v], u)
	}

	visited := make(map[int]bool)

	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}

		visited[i] = true
		for _, j := range graphMap[i] {
			dfs(j)
		}
	}

	res := 0
	for i := range n {
		if visited[i] {
			continue
		}

		res += 1
		dfs(i)
	}

	return res
}
