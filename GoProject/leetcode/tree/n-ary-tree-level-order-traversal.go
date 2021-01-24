package tree

func levelOrder(root *Node) (res [][]int) {
	var levelorder func(root []*Node)
	levelorder = func(roots []*Node) {
		var nodes []*Node
		var nodeVals []int

		if len(roots) == 0 {
			return
		}
		for _, node := range roots {
			if node != nil {
				nodeVals = append(nodeVals, node.Val)
				for _, n := range node.Children {
					if n != nil {
						nodes = append(nodes, n)
					}
				}
			}
		}
		if len(nodeVals) != 0 {
			res = append(res, nodeVals)
		}
		levelorder(nodes)
	}

	levelorder([]*Node{root})
	return
}
