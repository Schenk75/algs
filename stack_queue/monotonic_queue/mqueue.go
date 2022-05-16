package main

// 单调递减的队列
type MQueue struct {
	queue []int
}

func NewMQueue() *MQueue {
	return &MQueue{[]int{}}
}

// 删去队列中比x小的所有元素，然后将x添加到队尾
func (q *MQueue) Push(x int) {
	if !q.isEmpty() {
		idx := q.findPushIdx(x)		// 由于队列元素有序，可以使用二分搜索找到插入位置
		q.queue = q.queue[:idx]
	}
	q.queue = append(q.queue, x)
}

// 若队首元素等于x，则删去，否则不做任何操作
func (q *MQueue) Pop(x int) {
	if !q.isEmpty() && q.Front() == x {
		q.queue = q.queue[1:]
	}
}

func (q *MQueue) Front() int {
	return q.queue[0]
}

func (q *MQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *MQueue) findPushIdx(x int) int {
	l, r := 0, len(q.queue)-1
	for l <= r {
		mid := (l + r) / 2
		if q.queue[mid] < x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}