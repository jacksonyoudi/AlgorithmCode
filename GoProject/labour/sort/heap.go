package sort


func heapify(tree []int, n, i int) {
    c1 := 2 * i + 1
    c2 := 2 * i + 2

    max := i
    for c1 < n && tree[c1] > tree[i] {
        max = c1
    }


    for c2 < n && tree[c2] > tree[i] {
        max = c2
    }

    if max != i {
        tree[max],tree[i] = tree[i], tree[max]
        heapify(tree, n, max)
    }
}


func initHeap(tree []int, n int) {
    last_parent := (n - 1) / 2

    for i := last_parent; i >= 0; i-- {
        heapify(tree, n, i)



    }
}


func heap_sort(tree []int) {
    n := len(tree)
    initHeap(tree, n)
    for i := n -1; i >= 0 ;i-- {
        tree[i],tree[0] = tree[0], tree[i]
        heapify(tree, i, 0)
    }

}
