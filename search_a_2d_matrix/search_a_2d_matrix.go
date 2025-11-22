package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix[0])*len(matrix) - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])

		if matrix[row][col] > target {
			right = mid - 1
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
