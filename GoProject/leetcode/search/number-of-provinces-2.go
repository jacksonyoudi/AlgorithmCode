package main

func findCircleNum2(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)

	// 初始化 父节点执行自己
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	// 如果父节点不是自己就一直找先去， 并设置父节点
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// 合并
	union := func(from, to int) {

		// 可以考虑压缩
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return

}
