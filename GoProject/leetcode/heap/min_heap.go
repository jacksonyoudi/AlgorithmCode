package heap

type MinHeap struct {
	data    []int
	size    int
	maxSize int
}

func newMinHeap(maxSize int) *MinHeap {
	return &MinHeap{
		data:    []int{},
		size:    0,
		maxSize: maxSize,
	}
}

// 是否有叶子节点
func (m *MinHeap) leaf(index int) bool {
	if index >= m.size/2 && index <= m.size {
		return true
	}
	return false
}

func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

//func (m *MinHeap) Leftchild {
//
//}
