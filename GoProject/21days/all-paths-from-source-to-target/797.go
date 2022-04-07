package main

import "fmt"

// 深度遍历
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) - 1
	Data := &Data{
		[][]int{},
	}
	helper(graph, []int{0}, Data, 0, n)
	return Data.Path
}

func helper(graph [][]int, cur []int, path *Data, index int, n int) {
	ps := graph[index]
	for _, v := range ps {
		if v == n {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			tmp = append(tmp, v)
			path.Path = append(path.Path, tmp)
		} else {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			tmp = append(tmp, v)
			helper(graph, tmp, path, v, n)
		}
	}
}

type Data struct {
	Path [][]int
}

func main() {
	a := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	path := allPathsSourceTarget(a)
	fmt.Println(path)
}
