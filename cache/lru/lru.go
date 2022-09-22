package main

import (
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
	if len(arr) == 0 {
		return &NodeList{Len: 0}
	}
	dummy := NewNode(-1)
	prev := dummy
	var node *Node
	for _, val := range arr {
		node = NewNode(val)
		prev.Next = node
		node.Prev = prev
		prev = node
	}

	return &NodeList{
		Head: dummy.Next,
		Tail: node,
		Len:  len(arr),
	}
}

func (nl *NodeList) String() string {
	head := nl.Head
	var sb strings.Builder
	for i := 0; i < nl.Len; i++ {
		sb.WriteString(strconv.Itoa(head.Val))
		head = head.Next
		if i == nl.Len - 1 {break}
		sb.WriteString("->")
	}
	sb.WriteString(" len:")
	sb.WriteString(strconv.Itoa(nl.Len))
	return sb.String()
}

// 插入到链表头
func (nl *NodeList) Insert(val int) *Node {
	node := NewNode(val)
	if nl.Len == 0 {
		nl.Head = node
		nl.Tail = node
	} else {
		node.Next = nl.Head
		nl.Head.Prev = node
		nl.Head = node
	}
	nl.Len++
	return node
}

func (nl *NodeList) Evict() int {
	res := nl.Tail
	if nl.Len == 1 {
		nl.Head = nil
		nl.Tail = nil
	} else {
		nl.Tail = nl.Tail.Prev
		nl.Tail.Next = nil
	}
	res.Prev = nil
	nl.Len--
	return res.Val
}

func (nl *NodeList) Reinsert(node *Node) *Node {
	if node.Prev == nil {
		// 说明该节点本来就是head
		return node
	}
	node.Prev.Next = node.Next
	if node.Next == nil {
		// 该节点为尾节点
		nl.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nl.Head
	nl.Head.Prev = node
	nl.Head = node
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
