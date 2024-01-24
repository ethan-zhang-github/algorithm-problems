package linked_list

func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	i, j := headA, headB
	for i != j {
		if i == nil {
			i = headB
		} else {
			i = i.Next
		}
		if j == nil {
			j = headA
		} else {
			j = j.Next
		}
	}
	return i
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := findFromEnd(dummy, n+1)
	pre.Next = pre.Next.Next
	return dummy.Next
}

func findFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p, q, c := list1, list2, dummy
	for p != nil && q != nil {
		if p.Val <= q.Val {
			c.Next = p
			p = p.Next
		} else {
			c.Next = q
			q = q.Next
		}
		c = c.Next
	}
	if p != nil {
		c.Next = p
	}
	if q != nil {
		c.Next = q
	}
	return dummy.Next
}

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists, 0, len(lists))
}

func mergeKLists(lists []*ListNode, l int, r int) *ListNode {
	if r <= l {
		return nil
	}
	if r-l <= 1 {
		return lists[l]
	}
	mid := (l + r) / 2
	ll := mergeKLists(lists, l, mid)
	rl := mergeKLists(lists, mid, r)
	return MergeTwoLists(ll, rl)
}

func Partition(head *ListNode, x int) *ListNode {
	ll := &ListNode{Val: -1}
	rl := &ListNode{Val: -1}
	p, lp, rp := head, ll, rl
	for p != nil {
		if p.Val < x {
			lp.Next = p
			lp = lp.Next
		} else {
			rp.Next = p
			rp = rp.Next
		}
		p = p.Next
	}
	rp.Next = nil
	lp.Next = rl.Next
	return ll.Next
}
