package tree

import "fmt"

var NULL = -999

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(values []int) *TreeNode {
	return constructor(values, 0)
}

func constructor(values []int, i int) *TreeNode {
	if i >= len(values) || values[i] == NULL {
		return nil
	}
	root := TreeNode{Val: values[i]}
	root.Left = constructor(values, 2*i+1)
	root.Right = constructor(values, 2*i+2)
	return &root
}

func (t *TreeNode) size() int {
	return size(t)
}

func size(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return 1 + size(t.Left) + size(t.Right)
}

func (t *TreeNode) String() string {
	values := t.toArray()
	if len(values) == 0 {
		return "[]"
	}
	res := fmt.Sprintf("[%d", values[0])
	for i := 1; i < len(values); i++ {
		if values[i] == NULL {
			res += ", null"
		} else {
			res += fmt.Sprintf(", %d", values[i])
		}
	}
	res += "]"
	return res
}

func (t *TreeNode) toArray() []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		if pop == nil {
			res = append(res, NULL)
			continue
		}
		res = append(res, pop.Val)
		queue = append(queue, pop.Left, pop.Right)
	}
	return res
}
