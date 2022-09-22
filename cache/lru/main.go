package main

import "fmt"

func main() {
	// nlist := NewNodeList([]int{})
	// nlist.Insert(1)
	// nlist.Insert(2)
	// fmt.Println(nlist)
	testLRU()
	testLRU1()
	testLRU2()
	testLRU3()
}

func testLRU() {
	lru := Constructor(2)
	put(&lru, 1, 1)
	put(&lru, 2, 2)
	get(&lru, 1)  // 1
	put(&lru, 3, 3)
	get(&lru, 2)  // -1
	put(&lru, 4, 4)
	get(&lru, 1)  // -1
	get(&lru, 3)  // 3
	get(&lru, 4)  // 4
	fmt.Println()
}

func testLRU1() {
	lru := Constructor(1)
	put(&lru, 2, 1)
	get(&lru, 2)
	put(&lru, 3, 2)
	get(&lru, 2)
	get(&lru, 3)
	fmt.Println()
}

func testLRU2() {
	lru := Constructor(2)
	put(&lru, 2, 1)
	put(&lru, 2, 2)
	get(&lru, 2)
	put(&lru, 1, 1)
	put(&lru, 4, 1)
	get(&lru, 2)
	fmt.Println()
}

func testLRU3() {
	lru := Constructor(2)
	put(&lru, 2, 1)
	put(&lru, 1, 1)
	put(&lru, 2, 3)
	put(&lru, 4, 1)
	get(&lru, 1)
	get(&lru, 2)
	fmt.Println()
}

func put(lru *LRUCache, key, val int) {
	lru.Put(key, val)
	fmt.Printf("Put %d:%d ", key, val)
	fmt.Println(lru.KeyList)
}

func get(lru *LRUCache, key int) {
	fmt.Printf("Get %d ", lru.Get(key))
	fmt.Println(lru.KeyList)
}