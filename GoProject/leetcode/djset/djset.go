package main

import "os"

type DjSet struct {
	parent []int
	size   int
	rank   []int
}

func NewDjset(size int) *DjSet {
	d := &DjSet{
		parent: make([]int, size),
		rank:   make([]int, size),
		size:   size,
	}
	d.initial()
	return d

}

// 初始化，
func (d *DjSet) initial() {
	for i := 0; i < d.size; i++ {
		d.parent[i] = -1
		d.rank[i] = 0
	}
}

func (d *DjSet) findRoot(index int) int {
	root := index

	for root != -1 {
		root = d.parent[root]
	}
	return root
}

// 1 success 0 failure
func (d *DjSet) Union(x int, y int) int {
	xRoot := d.findRoot(x)
	yRoot := d.findRoot(y)

	if xRoot == yRoot {
		return 0
	}
	if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[yRoot] = xRoot
	} else if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[xRoot] = yRoot
	} else {
		d.parent[xRoot] = yRoot
		d.rank[yRoot] += 1
	}

	return 1
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
		is := djset.Union(x, y)
		if is == 0 {
			println("cycle detected!")
			os.Exit(1)
		}
	}

	println("no cycle")
}
