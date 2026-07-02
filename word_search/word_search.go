package word_search

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	visited := make([][]bool, rows)
	for i := range rows {
		visited[i] = make([]bool, cols)
	}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || r >= rows ||
			c < 0 || c >= cols ||
			word[i] != board[r][c] ||
			visited[r][c] == true {
			return false
		}

		visited[r][c] = true
		res := dfs(r-1, c, i+1) || // up
			dfs(r+1, c, i+1) || // down
			dfs(r, c-1, i+1) || // left
			dfs(r, c+1, i+1) // right
		visited[r][c] = false

		return res
	}

	for i := range rows {
		for j := range cols {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
