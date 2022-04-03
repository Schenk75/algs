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
	selectionSort(arr)
	fmt.Println(arr)
	sort.Ints(orig)
	fmt.Println(orig, "(correct ans)")
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i ++ {
		min := i
		for j := i+1; j < len(arr); j ++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		fmt.Println("current min:", arr[i], arr)
	}
}