package tree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tree := Constructor([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(maxDepth(tree))
}

func TestPreorderTraversal(t *testing.T) {
	tree := Constructor([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(preorderTraversal(tree))
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := Constructor([]int{1, NULL, 2, NULL, 3, 4, 5, 6, NULL, NULL, 7, 8, NULL, 9, NULL, 10, NULL, 11, NULL, 12, NULL, 13, NULL})
	fmt.Println(diameterOfBinaryTree(tree))
}

func TestFlatten(t *testing.T) {
	tree := Constructor([]int{1, 2, 5, 3, 4, NULL, 6})
	fmt.Println(tree)
	flatten(tree)
	fmt.Println(tree)
}
