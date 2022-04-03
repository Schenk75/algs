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
	arr = mergeSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 { return arr }
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i ++
		} else {
			res = append(res, right[j])
			j ++
		}
	}

	if i < len(left) {
		res = append(res, left[i:]...)
	}
	if j < len(right) {
		res = append(res, right[j:]...)
	}

	fmt.Println("merge:", left, right, "->", res)
	return res
}