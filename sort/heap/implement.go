package main

import "fmt"

// 小顶堆实现（要实现大顶堆只需要修改up和down中的比较即可）

// 自上而下堆化，用于建堆和删除元素
func down(heap *[]int, i int) {
	smallest := i
	lChild, rChild := 2*i+1, 2*i+2
	if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
		smallest = lChild
	}
	if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
		smallest = rChild
	}

	if i != smallest {
		(*heap)[smallest], (*heap)[i] = (*heap)[i], (*heap)[smallest]
		down(heap, smallest)
	}
}

// 自下而上堆化，用于插入元素
func up(heap *[]int, i int) {
	parent := (i-1) / 2
	if parent >= 0 && (*heap)[parent] > (*heap)[i] {
		(*heap)[parent], (*heap)[i] = (*heap)[i], (*heap)[parent]
		up(heap, parent)
	}
}

func buildHeap(heap []int) {
	n := len(heap)
	// 从非叶子节点开始，对自下而上的元素进行堆化
	start := n / 2 - 1
	for i := start; i >= 0; i-- {
		down(&heap, i)
	}
}

// 先将堆顶元素和最后的元素交换，再对堆顶元素进行一次堆化
func pop(heap *[]int) int {
	root := (*heap)[0]
	n := len(*heap)
	lastElement := (*heap)[n-1]
	(*heap)[0] = lastElement
	*heap = (*heap)[:n-1]
	down(heap, 0)
	return root
}

// 插入一个元素到最后，再向上堆化
func insert(heap *[]int, num int) {
	*heap = append((*heap), num)
	n := len(*heap)
	up(heap, n-1)
}

func heapRealize() {
	fmt.Println("\n=========================realize heap=========================")
	heap := []int{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	fmt.Println("arr: ", heap)
	buildHeap(heap)
	fmt.Println("heap: ", heap)
	fmt.Println("pop: ", pop(&heap))
	insert(&heap, 6)
	res := []int{}
	for len(heap) > 0 {
		res = append(res, pop(&heap))
	}
	fmt.Println("after sort: ", res)
}