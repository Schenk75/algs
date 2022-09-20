package main

import "fmt"

// 堆实现（小顶堆中less为true，大顶堆中less为false）
type MyHeap struct {
    heap []int
    less bool
}

func NewMyHeap(less bool) MyHeap {
    return MyHeap {
        heap: []int{},
        less: less,
    }
}

func (h *MyHeap) comp(a, b int) bool {
    if h.less {
        return a < b 
    } else {
        return a > b
    }
}

func (h *MyHeap) Len() int {
	return len(h.heap)
}

// 自上而下堆化，用于建堆和删除元素
func (h *MyHeap) down(i int) {
	target := i
	lChild, rChild := i << 1 + 1, i << 1 + 2
	if lChild < h.Len() && h.comp(h.heap[lChild], h.heap[target]) {
		target = lChild
	}
	if rChild < h.Len() && h.comp(h.heap[rChild], h.heap[target]) {
		target = rChild
	}

	if i != target {
		h.heap[target], h.heap[i] = h.heap[i], h.heap[target]
		h.down(target)
	}
}

// 自下而上堆化，用于插入元素
func (h *MyHeap) up(i int) {
	parent := (i-1) >> 1
	if parent >= 0 && h.comp(h.heap[i], h.heap[parent]) {
		h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		h.up(parent)
	}
}

func (h *MyHeap) BuildHeap(arr []int) {
	n := len(arr)
	h.heap = make([]int, n)
	copy(h.heap, arr)
	// 从非叶子节点开始，对自下而上的元素进行堆化
	start := n / 2 - 1
	for i := start; i >= 0; i-- {
		h.down(i)
	}
}

// 先将堆顶元素和最后的元素交换，再对堆顶元素进行一次堆化
func (h *MyHeap) Pop() int {
	root := h.heap[0]
	n := h.Len()
	lastElement := h.heap[n-1]
	h.heap[0] = lastElement
	h.heap = h.heap[:n-1]
	h.down(0)
	return root
}

// 插入一个元素到最后，再向上堆化
func (h *MyHeap) Insert(num int) {
	h.heap = append(h.heap, num)
	n := h.Len()
	h.up(n-1)
}

func heapRealize() {
	fmt.Println("\n=========================realize heap=========================")
	arr := []int{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	fmt.Println("arr: ", arr)

	// 小顶堆
	heap := NewMyHeap(true)
	// // 大顶堆
	// heap := NewMyHeap(false)

	heap.BuildHeap(arr)
	fmt.Println("heap: ", heap)
	fmt.Println("pop: ", heap.Pop())
	heap.Insert(6)
	res := []int{}
	for heap.Len() > 0 {
		res = append(res, heap.Pop())
	}
	fmt.Println("after sort: ", res)
}