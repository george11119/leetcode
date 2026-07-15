package number_of_islands

type Pair struct {
	row int
	col int
}

func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	visited := make(map[Pair]bool)

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rowLen || c >= colLen || grid[r][c] == '0' {
			return
		}

		if visited[Pair{r, c}] {
			return
		}

		visited[Pair{r, c}] = true

		for _, d := range directions {
			dfs(r+d[0], c+d[1])
		}
	}

	res := 0
	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if !visited[Pair{r, c}] && grid[r][c] != '0' {
				dfs(r, c)
				res++
			}
		}
	}

	return res
}
