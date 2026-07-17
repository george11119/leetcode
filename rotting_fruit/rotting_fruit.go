package rotting_fruit

type Pair struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	row, col := len(grid), len(grid[0])
	visited := make(map[Pair]bool)

	q := make([]Pair, 0)
	for r := range row {
		for c := range col {
			if grid[r][c] == 2 {
				q = append(q, Pair{r, c})
				visited[Pair{r, c}] = true
			}
		}
	}

	res := 0
	for len(q) != 0 {
		qLen := len(q)
		for qLen != 0 {
			cur := q[0]
			q = q[1:]

			for _, d := range directions {
				r := cur.row + d[0]
				c := cur.col + d[1]

				if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] != 1 {
					continue
				}

				if _, exists := visited[Pair{r, c}]; !exists {
					grid[r][c] = 2
					q = append(q, Pair{r, c})
					visited[Pair{r, c}] = true
				}
			}

			qLen--
		}

		if len(q) != 0 {
			res++
		}
	}

	for r := range row {
		for c := range col {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return res
}
