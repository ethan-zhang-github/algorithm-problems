package num_array

type Item struct {
	val   int
	index int
}

type PriorityQueue []*Item

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].val > q[j].val
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*q = old[0 : n-1]
	return item
}
