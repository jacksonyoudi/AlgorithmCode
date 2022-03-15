package main

type Heap struct {
	data []int
}

// 获取最大值
func (h *Heap) top() int {
	return h.data[0]
}

func (h *Heap) Push(i int) {
	// 先放到最后一位，然后上浮
	h.data = append(h.data, i)
	h.swim(len(h.data) - 1)

}


func (h *Heap) pop() int {
	m := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.sink(0)
	return m
}

func (h *Heap) swim(pos int) {
	for pos > 1 && h.data[pos/2] < h.data[pos] {
		h.data[pos/2], h.data[pos] = h.data[pos], h.data[pos/2]
		pos = pos / 2
	}
}

func (h *Heap) sink(pos int) {
	n := len(h.data)

	for 2*pos <= n {
		i := 2 * pos
		if i < n && h.data[i] < h.data[i+1] {
			i++
		}

		if h.data[pos] > h.data[i] {
			break
		}
		h.data[pos], h.data[i] = h.data[i], h.data[pos]
		pos = i
	}
}
