package tree

type DjSet struct {
    size int
    rank []int
    parent []int
}

func (d *DjSet) init() {
    for i := 0; i < d.size; i++ {
        d.rank[i] = 0
        d.parent[i] = i

    }
}

func GenDjSet(size int) *DjSet {
    return &DjSet{
        size: size,
        rank: make([]int, size),
        parent: make([]int, size),

    }
}


func (d *DjSet) Find(index int) int  {
    if index >= d.size {
        return -1
    }
    root := i
    if d.parent[root] != root {
        root = d.parent[i]
    }
    return root
}


func (d *DjSet) Union (x, y int)  {
    xRoot := d.findRoot(x)
    yRoot := d.findRoot(y){
    if xRoot == yRoot {
        return
    }
    if d.rank[xRoot] > d.rank[yRoot] {
        d.parent[xRoot] = d.rank[yRoot]
    }





}
