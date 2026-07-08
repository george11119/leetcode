package n_queens

func isSafe(board [][]bool, row int, col int) bool {
	// vertical check
	for c := 0; c < len(board); c++ {
		if board[row][c] {
			return false
		}
	}

	// horizontal check
	for r := 0; r < len(board); r++ {
		if board[r][col] {
			return false
		}
	}

	// diagonal checks
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] {
			return false
		}
	}
	for r, c := row+1, col+1; r < len(board) && c < len(board); r, c = r+1, c+1 {
		if board[r][c] {
			return false
		}
	}

	for r, c := row-1, col+1; r >= 0 && c < len(board); r, c = r-1, c+1 {
		if board[r][c] {
			return false
		}
	}
	for r, c := row+1, col-1; r < len(board) && c >= 0; r, c = r+1, c-1 {
		if board[r][c] {
			return false
		}
	}

	return true
}

func convertToStringSlice(board [][]bool) []string {
	res := make([]string, len(board))
	for i, r := range board {
		inner := make([]byte, 0)

		for j := range r {
			if board[i][j] {
				inner = append(inner, 'Q')
			} else {
				inner = append(inner, '.')
			}
		}

		res[i] = string(inner)
	}

	return res
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0, n)

	board := make([][]bool, 0, n)
	for range n {
		board = append(board, make([]bool, n))
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(board) {
			res = append(res, convertToStringSlice(board))
			return
		}

		for j := range board {
			if isSafe(board, i, j) {
				board[i][j] = true
				dfs(i + 1)
				board[i][j] = false
			}
		}
	}
	dfs(0)

	return res
}
