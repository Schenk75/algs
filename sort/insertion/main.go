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
	insertionSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i ++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j --
		}
		arr[j+1] = tmp
		fmt.Println("insert:", tmp, arr)
	}
}