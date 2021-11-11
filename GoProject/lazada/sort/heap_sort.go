package main

import "fmt"

func heapify(tree []int, n, i int) {
	c1 := 2*i + 1
	c2 := 2*i + 2
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

func InitHeap(tree []int, n int) {
	lastparent := (n - 1) / 2

	for i := lastparent; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

func heap_sort(tree []int) {
	n := len(tree)
	InitHeap(tree, n)

	for i := n - 1; i >= 0; i-- {
		tree[i], tree[0] = tree[0], tree[i]
		heapify(tree, i, 0)
	}

}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	heap_sort(arr)
	fmt.Println(arr)
}
