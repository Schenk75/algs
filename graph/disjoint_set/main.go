package main

import "fmt"

func main() {
	// eg.根据边的集合计算连通分量的数目
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}
	n := 6 // 总顶点数
	fmt.Println(n, edges)
	fmt.Println("count sets:", countSet(edges, n))
}

type unionFind struct {
	parent, size []int
	setCount     int // 当前连通分量数目
}

// 创建一个n节点的并查集
func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	parentX, parentY := uf.find(x), uf.find(y)
	if parentX == parentY { return }
	// 把小的集合合并到大的集合里去
	if uf.size[x] < uf.size[y] {
		parentX, parentY = parentY, parentX
	}
	uf.parent[parentY] = parentX
	uf.size[parentX] += uf.size[parentY]
	uf.setCount --	// 合并两个连通分量之后，连通分量数-1
}

func countSet(edges [][]int, n int) int {
	uf := newUnionFind(n)
	for _, edge := range edges {
		uf.union(edge[0], edge[1])
	}
	return uf.setCount
}