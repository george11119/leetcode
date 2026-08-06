package rotting_fruit

type Pair struct {
	r int
	c int
}

func orangesRotting(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	res := 0
	numOranges := 0

	q := make([]Pair, 0)
	visited := make(map[Pair]bool)

	for i := range row {
		for j := range col {
			orange := grid[i][j]
			if orange == 2 {
				visited[Pair{i,j}] = true
				q = append(q, Pair{i, j})
			}

			if orange == 1 {
				numOranges++
			}
		}
	}

	for len(q) != 0 {
		lenQ := len(q)

		for lenQ != 0 {
			coord := q[0]
			q = q[1:]

			for _, d := range directions {
				r := coord.r + d[0]
				c := coord.c + d[1]
				if r < 0 || r >= row || c < 0 || c >= col || grid[r][c] != 1 {
					continue
				}

				if exists := visited[Pair{r,c}]; !exists {
					visited[Pair{r,c}] = true
					q = append(q, Pair{r, c})
					grid[r][c] = 2
					numOranges--
				}
			}

			lenQ--
		}

		if len(q) != 0 {
			res++
		}
	}

	if numOranges != 0 {
		return -1
	}

	return res
}
