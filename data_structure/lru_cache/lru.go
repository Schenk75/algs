package main

import (
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

type NodeList struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewNodeList(arr []int) *NodeList {
	dummyHead := NewNode(-1)
	dummyTail := NewNode(-1)
	prev := dummyHead
	node := dummyHead
	for _, val := range arr {
		node = NewNode(val)
		prev.Next = node
		node.Prev = prev
		prev = prev.Next
	}
	node.Next = dummyTail
	dummyTail.Prev = node
	return &NodeList{
		Head: dummyHead,
		Tail: dummyTail,
		Len: len(arr),
	}
}

func (nl *NodeList) String() string {
	node := nl.Head.Next
	var sb strings.Builder
	for node != nl.Tail {
		sb.WriteString(strconv.Itoa(node.Val))
		node = node.Next
		if node == nl.Tail {break}
		sb.WriteString("->")
	}
	sb.WriteString(" len:")
	sb.WriteString(strconv.Itoa(nl.Len))
	return sb.String()
}

// 插入到链表头
func (nl *NodeList) Insert(val int) *Node {
	node := NewNode(val)
	next := nl.Head.Next
	nl.Head.Next = node
	node.Prev = nl.Head
	node.Next = next
	next.Prev = node
	nl.Len ++
	return node
}

func (nl *NodeList) Evict() int {
	if nl.Len == 0 {
		return math.MaxInt
	}
	res := nl.Tail.Prev
	res.Prev.Next = nl.Tail
	nl.Tail.Prev = res.Prev
	nl.Len --
	return res.Val
}

func (nl *NodeList) Reinsert(node *Node) *Node {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev = nl.Head
	node.Next = nl.Head.Next
	nl.Head.Next = node
	node.Next.Prev = node
	return node
}

type LRUVal struct {
	Val  int
	node *Node
}

type LRUCache struct {
	KeyList  *NodeList
	KVMap    map[int]*LRUVal
	Capacity int
}

func Constructor(capacity int) LRUCache {
	mp := make(map[int]*LRUVal, capacity)
	return LRUCache{
		KeyList:  NewNodeList([]int{}),
		KVMap:    mp,
		Capacity: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if lruVal, ok := lru.KVMap[key]; ok {
		newNode := lru.KeyList.Reinsert(lruVal.node)
		lruVal.node = newNode
		return lruVal.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lruVal, ok := lru.KVMap[key]; ok {
		lruVal.Val = value
		// 插入新值也要更新LRU
		newNode := lru.KeyList.Reinsert(lruVal.node)
		lruVal.node = newNode
	} else {
		if lru.KeyList.Len == lru.Capacity {
			// 需要先删除最久未使用元素
			evictKey := lru.KeyList.Evict()
			delete(lru.KVMap, evictKey)
		}
		node := lru.KeyList.Insert(key)
		lru.KVMap[key] = &LRUVal{Val: value, node: node}
	}
}
