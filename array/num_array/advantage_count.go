package num_array

import (
	"container/heap"
	"sort"
)

// AdvantageCount https://leetcode.cn/problems/advantage-shuffle/
func AdvantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i, v := range nums2 {
		item := Item{val: v, index: i}
		heap.Push(&pq, &item)
	}
	res := make([]int, len(nums1))
	l, r := 0, len(nums1)-1
	for pq.Len() > 0 {
		top := heap.Pop(&pq).(*Item)
		if nums1[r] > top.val {
			res[top.index] = nums1[r]
			r--
		} else {
			res[top.index] = nums1[l]
			l++
		}
	}
	return res
}
