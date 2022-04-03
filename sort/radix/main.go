package main

import (
	"fmt"
	"math"
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
	radixSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func radixSort(arr []int) {
	maxV := max(arr)
	k := 0	// 最大位数，即需要排序几趟
	for maxV > 0 {
		k ++
		maxV /= 10
	}

	for i := 0; i < k; i ++ {
		buckets := make([][]int, 10)
		for _, num := range arr {
			idx := getKthBit(num, i)
			buckets[idx] = append(buckets[idx], num)
		}
		fmt.Println("kth bit:", i, buckets)

		idx := 0
		for j := range buckets {
			for len(buckets[j]) > 0 {
				arr[idx] = buckets[j][0]
				buckets[j] = buckets[j][1:]
				idx ++
			}
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

func getKthBit(num int, k int) int {
	num = num / int(math.Pow10(k))
	return num % 10
}