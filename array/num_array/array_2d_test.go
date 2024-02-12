package num_array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{7, 6, 5, 4},
		{4, 1, 2, 3},
	}
	Rotate(matrix)
	fmt.Println(matrix)
}

func TestSpiralArray(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{7, 6, 5, 4},
		{4, 1, 2, 3},
	}
	fmt.Println(SpiralArray(matrix))
}
