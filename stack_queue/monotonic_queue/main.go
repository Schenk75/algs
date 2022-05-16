package main

import "fmt"

// leetcode239 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。 返回滑动窗口中的最大值。
func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
    q := NewMQueue()
    res := []int{}
    for i := 0; i < k; i++ {
        q.Push(nums[i])
    }
    res = append(res, q.Front())

    for i := k; i < len(nums); i++ {
        q.Pop(nums[i-k])
        q.Push(nums[i])
        res = append(res, q.Front())
    }
    return res
}