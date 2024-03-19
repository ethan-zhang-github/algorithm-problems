package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalRecursive(root, make([]int, 0))
}

func preorderTraversalRecursive(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}
	values = append(values, root.Val)
	values = preorderTraversalRecursive(root.Left, values)
	values = preorderTraversalRecursive(root.Right, values)
	return values
}
