package max_area_of_island

type Pair struct {
	row int
	col int
}

func maxAreaOfIsland(grid [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	row := len(grid)
	col := len(grid[0])
	visited := make(map[Pair]bool)

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= row || c >= col || visited[Pair{r, c}] || grid[r][c] == 0 {
			return 0
		}

		res := 0

		if grid[r][c] == 1 {
			res += 1
			visited[Pair{r, c}] = true
		}

		for _, d := range directions {
			res += dfs(r+d[0], c+d[1])
		}

		return res
	}

	let := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 && !visited[Pair{r, c}] {
				let = max(dfs(r, c), let)
			}
		}
	}

	return let
}
