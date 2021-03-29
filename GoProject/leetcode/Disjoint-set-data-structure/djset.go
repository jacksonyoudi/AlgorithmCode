package main

import "os"

type DjSet struct {
	parent []int
	rank   []int
	size   int
}

func NewDjset(size int) *DjSet {
	d := &DjSet{
		parent: make([]int, size),
		rank:   make([]int, size),
		size:   size,
	}
	d.initialise()
	return d
}

func (d *DjSet) initialise() {
	for i := 0; i < d.size; i++ {
		d.parent[i] = -1
		d.rank[i] = 0
	}
}

func (d *DjSet) findRoot(index int) int {
	// 默认根节点就是自己
	root := index
	for d.parent[root] != -1 {
		root = d.parent[root]
	}
	return root
}

// 1 success 0 failure
func (d *DjSet) unionVertices(x int, y int) int {
	rootX := d.findRoot(x)
	rootY := d.findRoot(y)

	if rootX == rootY {
		return 0
	} else {
		if d.rank[rootX] > d.rank[rootX] {
			d.parent[rootY] = rootX
		} else if d.rank[rootX] < d.rank[rootX] {
			d.parent[rootX] = rootY
		} else {
			d.parent[rootX] = rootY
			d.rank[rootY] += 1
		}
		return 1
	}

}

func main() {

	edges := [][]int{
		{0, 1},
		{1, 2},
		{1, 3},
		{3, 4},
		{5, 4},
		{2, 5},
	}

	djset := NewDjset(6)

	for _, items := range edges {
		x := items[0]
		y := items[1]
		is := djset.unionVertices(x, y)
		if is == 0 {
			println("cycle detected!")
			os.Exit(1)
		}
	}

	println("no cycle")

}
