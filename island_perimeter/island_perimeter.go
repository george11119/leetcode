package island_perimeter

func islandPerimeter(grid [][]int) int {
	// create a visited array so I dont create cycles
	visited := make([][]int, len(grid))
	for i := range len(visited) {
		visited[i] = make([]int, len(grid[0]))
	}

	res := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == 0 {
			res += 1
			return
		}

		if visited[r][c] == 1 {
			return
		}

		visited[r][c] = 1

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				return res
			}
		}
	}

	return 0
}
