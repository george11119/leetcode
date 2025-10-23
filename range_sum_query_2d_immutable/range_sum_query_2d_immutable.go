package range_sum_query_2d_immutable

// Calculate range sum on the fly
// time complexity: O(n * m)
// space complexity: O(1)
//type NumMatrix struct {
//	matrix [][]int
//}
//
//func Constructor(matrix [][]int) NumMatrix {
//	return NumMatrix{matrix}
//}
//
//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//	sum := 0
//	for i := row1; i <= row2; i++ {
//		for j := col1; j <= col2; j++ {
//			sum += this.matrix[i][j]
//		}
//	}
//	return sum
//}

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	sumMat := make([][]int, row+1)
	for i := 0; i < len(sumMat); i++ {
		sumMat[i] = make([]int, col+1)
	}

	for r := 0; r < row; r++ {
		sum := 0
		for c := 0; c < col; c++ {
			sum += matrix[r][c]
			above := sumMat[r][c+1]
			sumMat[r+1][c+1] = sum + above
		}
	}

	return NumMatrix{sumMat}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	topLeft := this.sumMatrix[row1][col1]
	topRight := this.sumMatrix[row1][col2+1]
	botLeft := this.sumMatrix[row2+1][col1]
	botRight := this.sumMatrix[row2+1][col2+1]

	return botRight + topLeft - topRight - botLeft
}
