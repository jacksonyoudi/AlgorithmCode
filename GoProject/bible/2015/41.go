package _015

type Node struct {
	Data int
	Next *Node
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func deleteAbsNode(root *Node) *Node {
	m := make(map[int]int, 0)

	dummpy := root

	pre := dummpy

	for root != nil {
		val := root.Data
		if _, ok := m[abs(val)]; ok {
			pre.Next = root.Next
		} else {
			m[abs(val)] = 1
		}
		pre = pre.Next
		root = root.Next
	}

	return dummpy.Next
}
