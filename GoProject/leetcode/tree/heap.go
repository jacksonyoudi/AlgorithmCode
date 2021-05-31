package tree

func heapify(tree []int, n, i int) {
	c1 := 2*i + 1
	c2 := 2*i + 2

	// 判断 是否出界
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max)

	}

}

func initHeap(tree []int, n int) {
	lastParent := (n - 1) / 2

	for i := lastParent; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

func heap_sort(tree []int, n int) {
	initHeap(tree, n)

	for i := n - 1; i >= 0; i-- {
		tree[i], tree[0] = tree[0], tree[i]
		heapify(tree, i, 0)
	}

}
