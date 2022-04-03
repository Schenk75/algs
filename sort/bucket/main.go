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
	bucketSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func bucketSort(arr []int) {
	//桶数
	n := len(arr)
	max := max(arr)
	buckets := make([][]int, n)

	index := 0
	for i := 0; i < n; i++ {
		index = arr[i] * (n-1) /max	//分配桶index = value * (n-1) /k
		buckets[index] = append(buckets[index], arr[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < n; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			sort.Ints(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
			fmt.Println("bucket:", i, buckets[i])
		}
	}
}

func max(arr []int) int {
	res := arr[0]
	for _, n := range arr {
		if n > res {
			res = n
		}
	}
	return res
}