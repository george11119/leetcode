package valid_sudoku

func IsValidSudoku(board [][]byte) bool {
	rowCheck := make([]map[byte]bool, 9)
	colCheck := make([]map[byte]bool, 9)
	squareCheck := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowCheck[i] = make(map[byte]bool)
		colCheck[i] = make(map[byte]bool)
		squareCheck[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			boardVal := board[r][c]
			if boardVal == '.' {
				continue
			}

			_, inRow := rowCheck[r][boardVal]
			_, inCol := colCheck[c][boardVal]
			squareIdx := (r/3)*3 + c/3
			_, inSquare := squareCheck[squareIdx][boardVal]

			// this means there is a duplicate value
			if inRow || inCol || inSquare {
				return false
			} else {
				rowCheck[r][boardVal] = true
				colCheck[c][boardVal] = true
				squareCheck[squareIdx][boardVal] = true
			}
		}
	}

	return true
}
