package surrounded_regions

type Pair struct {
	row int
	col int
}

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make(map[Pair]bool)

	q := make([]Pair, 0)

	for r := range rows {
		if board[r][0] == 'O' {
			q = append(q, Pair{r, 0})
		}
		if board[r][cols-1] == 'O' {
			q = append(q, Pair{r, cols - 1})
		}
	}

	for c := range cols {
		if board[0][c] == 'O' {
			q = append(q, Pair{0, c})
		}
		if board[rows-1][c] == 'O' {
			q = append(q, Pair{rows - 1, c})
		}
	}

	for len(q) != 0 {
		coord := q[0]
		q = q[1:]

		for _, d := range directions {
			r := coord.row + d[0]
			c := coord.col + d[1]

			if r < 0 || r >= rows || c < 0 || c >= cols || visited[Pair{r, c}] {
				continue
			}

			visited[coord] = true
			if board[r][c] == 'O' {
				q = append(q, Pair{r, c})
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if board[r][c] == 'O' && !visited[Pair{r, c}] {
				board[r][c] = 'X'
			}
		}
	}
}
