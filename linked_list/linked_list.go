package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConsCycle(cycleIndex int, nums ...int) *ListNode {
	l := ConsWithArray(nums)
	if cycleIndex >= len(nums) {
		return l
	}
	l.Tail().Next = l.Get(cycleIndex)
	return l
}

func Cons(nums ...int) *ListNode {
	return ConsWithArray(nums)
}

func ConsWithArray(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	head.Next = ConsWithArray(nums[1:])
	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	res := fmt.Sprintf("%d", l.Val)
	visited := make(map[*ListNode]int)
	visited[l] = 0
	for index, node := 1, l.Next; node != nil; index, node = index+1, node.Next {
		if val, ok := visited[node]; ok {
			res += fmt.Sprintf(" -> Index(%d)", val)
			return res
		}
		res += fmt.Sprintf(" -> %d", node.Val)
		visited[node] = index
	}
	res += " -> nil"
	return res
}

func (l *ListNode) Len() int {
	if l == nil {
		return 0
	}
	return 1 + l.Next.Len()
}

func (l *ListNode) Tail() *ListNode {
	if l == nil {
		return nil
	}
	if l.Next == nil {
		return l
	}
	return l.Next.Tail()
}

func (l *ListNode) Get(index int) *ListNode {
	if l == nil {
		return nil
	}
	if index == 0 {
		return l
	}
	return l.Next.Get(index - 1)
}

func (l *ListNode) Concat(other *ListNode) *ListNode {
	tail := l.Tail()
	if tail == nil {
		return other
	}
	tail.Next = other
	return l
}
