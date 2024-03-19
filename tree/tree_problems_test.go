package tree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tree := Constructor([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(maxDepth(tree))
}
