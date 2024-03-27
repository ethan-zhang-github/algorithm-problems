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

// https://leetcode.cn/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	traverse(root, func(node *TreeNode) {
		maxDiameter = max(maxDiameter, maxDepth(node.Left)+maxDepth(node.Right))
	})
	return maxDiameter
}

func traverse(root *TreeNode, f func(node *TreeNode)) {
	if root == nil {
		return
	}
	f(root)
	traverse(root.Left, f)
	traverse(root.Right, f)
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	flatten(root.Left)
	l := root.Left
	for l.Right != nil {
		l = l.Right
	}

	flatten(root.Right)
	l.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		var pre *Node
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int, len(inorder))
	for v, i := range inorder {
		indexMap[v] = i
	}
	return buildTreeRecursive(preorder, 0, len(preorder), inorder, 0, len(inorder), indexMap)
}

func buildTreeRecursive(preorder []int, a int, b int, inorder []int, c int, d int, indexMap map[int]int) *TreeNode {
	
	root := preorder[a]
	rootIndex := indexMap[root]
	left := buildTreeRecursive(preorder, a+1, a+1+rootIndex-c, inorder, c, rootIndex, indexMap)
	//left := buildTreeRecursive(preorder[1:1+rootIndex], inorder[0:rootIndex], indexMap)
	right := buildTreeRecursive(preorder, a+2+rootIndex-c, b, inorder, rootIndex+1, d, indexMap)
	//right := buildTreeRecursive(preorder[rootIndex+1:], inorder[rootIndex+1:], indexMap)
	return &TreeNode{Val: root, Left: left, Right: right}
}
