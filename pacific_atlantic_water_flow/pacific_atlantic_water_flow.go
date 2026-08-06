package pacific_atlantic_water_flow

type Pair struct {
	row int
	col int
}

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := make([][]int, 0)

	pacific := make(map[Pair]bool)
	atlantic := make(map[Pair]bool)

	pacificQueue := make([]Pair, 0)
	atlanticQueue := make([]Pair, 0)

	for i := range rows {
		pacificQueue = append(pacificQueue, Pair{i, 0})
		atlanticQueue = append(atlanticQueue, Pair{i, cols - 1})
	}

	for j := range cols {
		pacificQueue = append(pacificQueue, Pair{0, j})
		atlanticQueue = append(atlanticQueue, Pair{rows - 1, j})
	}

	bfs := func(q []Pair, trackerMap map[Pair]bool) {
		visited := make(map[Pair]bool)

		for len(q) != 0 {
			coord := q[0]
			q = q[1:]

			trackerMap[Pair{coord.row, coord.col}] = true

			for _, d := range directions {
				r := coord.row + d[0]
				c := coord.col + d[1]

				if r < 0 || r >= rows || c < 0 || c >= cols || visited[Pair{r, c}] {
					continue
				}

				if heights[coord.row][coord.col] <= heights[r][c] {
					q = append(q, Pair{r, c})
					visited[Pair{r, c}] = true
				}
			}
		}
	}

	bfs(pacificQueue, pacific)
	bfs(atlanticQueue, atlantic)

	for i := range rows {
		for j := range cols {
			if pacific[Pair{i, j}] && atlantic[Pair{i, j}] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
