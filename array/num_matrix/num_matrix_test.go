package num_matrix

import "testing"

func TestSumRegion(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	if ans := numMatrix.SumRegion(2, 1, 4, 3); ans != 8 {
		t.Errorf("SumRegion(2, 1, 4, 3) expected be 8, but %d got", ans)
	}
}
