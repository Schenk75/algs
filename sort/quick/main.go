package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 10
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(101))
	}
	orig := append([]int{}, arr...)
	fmt.Println("origin arr:", orig)
	quickSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func quickSort(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left, right int) {
	if left >= right { return }
	pivot := left
	idx := pivot + 1	// 存放最后一个比arr[pivot]小的元素的下标
	for i := idx; i <= right; i ++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[idx] = arr[idx], arr[i]
			idx ++
		}
	}
	arr[pivot], arr[idx-1] = arr[idx-1], arr[pivot]
	fmt.Println("pivot:", arr[idx-1], arr)
	quick(arr, pivot, idx-1)
	quick(arr, idx, right)
}