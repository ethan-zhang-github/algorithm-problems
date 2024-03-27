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

func TestConnect(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5}},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7}},
	}
	connect(root)
}

func TestBuildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preOrder, inOrder)
	fmt.Println(tree)
}
