package num_matrix

// NumMatrix https://leetcode.cn/problems/range-sum-query-2d-immutable/description/
type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	h, w := len(matrix), len(matrix[0])
	sumMatrix := make([][]int, h)
	for i := 0; i < h; i++ {
		sumMatrix[i] = make([]int, w)
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				sumMatrix[i][j] = matrix[i][j]
			} else if i == 0 {
				sumMatrix[i][j] = sumMatrix[i][j-1] + matrix[i][j]
			} else if j == 0 {
				sumMatrix[i][j] = sumMatrix[i-1][j] + matrix[i][j]
			} else {
				sumMatrix[i][j] = sumMatrix[i-1][j] + sumMatrix[i][j-1] - sumMatrix[i-1][j-1] + matrix[i][j]
			}
		}
	}
	return NumMatrix{sumMatrix: sumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sumMatrix[row2][col2]
	} else if row1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row2][col1-1]
	} else if col1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row1-1][col2]
	} else {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row2][col1-1] - this.sumMatrix[row1-1][col2] + this.sumMatrix[row1-1][col1-1]
	}
}
