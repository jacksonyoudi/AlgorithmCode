package tree

type DjSet struct {
	parent []int
	rank   []int
	size   int
}

func (d *DjSet) init() {
	for i := 0; i < d.size; i++ {
		d.parent[i] = i
		d.rank[i] = 0
	}
}

func (d *DjSet) Find(index int) int {
	root := d.parent[index]
	return root
}

func (d *DjSet) Union(x, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)

	if yRoot == xRoot {
		return
	}

	if d.rank[xRoot] > d.rank[yRoot] {
		d.parent[yRoot] = xRoot
	} else if d.rank[xRoot] < d.rank[yRoot] {
		d.parent[xRoot] = yRoot
	} else {
		d.parent[xRoot] = yRoot
		d.rank[yRoot] += 1
	}

}
