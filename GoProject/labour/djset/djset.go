package djset

type DjSet struct {
	// 父节点
	parent []int
	// 长度
	size int
	// 层数
	rank []int
}

func NewDjset(size int) *DjSet {
	d := &DjSet{
		parent: make([]int, size),
		rank:   make([]int, size),
		size:   size,
	}

	d.init()
	return d
}

func (d *DjSet) init() {
	for i := 0; i < d.size; i++ {
		d.parent[i] = -1
		d.rank[i] = 0
	}
}

// 如果没有父节点， 就是自身
func (d *DjSet) findRoot(index int) int {
	root := index

	v := d.parent[index]
	if v != -1 {
		root = v
	}
	return root
}

//  1 success 0 failure
func (d *DjSet) Union(x, y int) int {
	xRoot := d.findRoot(x)
	yRoot := d.findRoot(y)

	// x y 是一个父节点
	if xRoot == yRoot {
		return 0
	}

	if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[yRoot] = xRoot
	} else if d.rank[xRoot] < d.rank[yRoot] {
		d.parent[xRoot] = yRoot
	} else {
		d.parent[yRoot] = xRoot
		d.rank[xRoot] += 1
	}
	return 1

}
