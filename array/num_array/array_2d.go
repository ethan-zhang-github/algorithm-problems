package num_array

func Rotate(matrix [][]int) {
	for i := 0; i < len(matrix)-1; i++ {
		for j := i + 1; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	for _, row := range matrix {
		reverse(row)
	}
}

func reverse(array []int) {
	i, j := 0, len(array)-1
	for i < j {
		temp := array[i]
		array[i] = array[j]
		array[j] = temp
		i++
		j--
	}
}

// SpiralArray https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/description/
func SpiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	rows, columns := len(array), len(array[0])
	top, bottom, left, right := 0, rows-1, 0, columns-1
	res := make([]int, rows*columns)
	index := 0
	for top <= bottom && left <= right {
		for column := left; column <= right; column++ {
			res[index] = array[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			res[index] = array[row][right]
			index++
		}
		if top < bottom && left < right {
			for column := right - 1; column >= left; column-- {
				res[index] = array[bottom][column]
				index++
			}
			for row := bottom - 1; row > top; row-- {
				res[index] = array[row][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return res
}
