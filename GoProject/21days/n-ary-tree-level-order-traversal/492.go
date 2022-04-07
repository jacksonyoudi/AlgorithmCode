package n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

type Result struct {
	Data [][]int
}

func bfs(nodes []*Node, result *Result) {
	if len(nodes) == 0 {
		return
	}

	mNodes := make([]*Node, 0)
	mResult := make([]int, 0)
	for _, node := range nodes {
		if node != nil {
			mResult = append(mResult, node.Val)

			for _, child := range node.Children {
				mNodes = append(mNodes, child)
			}
		}
	}

	result.Data = append(result.Data, mResult)
	bfs(mNodes, result)
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := &Result{
		Data: make([][]int, 0),
	}
	bfs([]*Node{root}, result)
	return result.Data
}
