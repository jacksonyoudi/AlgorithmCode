package tree

func symmetry(nodes []*TreeNode) bool {
	res := []int{}
	rets := []*TreeNode{}

	if len(nodes)%2 != 0 && len(nodes) != 1 {
		return false
	}
	for _, node := range nodes {
		res = append(res, node.Val)
		if node != nil {
			if node.Left != nil {
				rets = append(rets, node.Left)
			}
			if node.Right != nil {
				rets = append(rets, node.Right)
			}
		}
	}

	if len(res)%2 != 0 && len(res) != 1 {
		return false
	}

	for i := 0; i < len(res)/2; i++ {
		if res[i] != res[len(res)-1-i] {
			return false
		}
	}
	if len(rets) == 0 {
		return true
	}

	return symmetry(rets)
}

func isSymmetric(root *TreeNode) bool {
	return symmetry([]*TreeNode{root})
}
