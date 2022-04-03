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
	countSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func countSort(arr []int) {
	max, min := findMaxMin(arr)
	cnt := make([]int, max-min+1)
	for _, num := range arr {
		cnt[num-min] ++
	}
	fmt.Println("cnt:", cnt)
	idx := 0
	for i := range cnt {
		for cnt[i] > 0 {
			arr[idx] = i + min
			cnt[i] --
			idx ++
		}
	}
}

func findMaxMin(arr []int) (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for _, n := range arr {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max, min
}