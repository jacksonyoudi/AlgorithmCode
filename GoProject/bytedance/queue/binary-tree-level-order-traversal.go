package queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type levelOrderData struct {
	Data [][]int
}

func levelOrderHelper(nodes []*TreeNode, result *levelOrderData) {
	if len(nodes) == 0 {
		return
	}

	newNodes := []*TreeNode{}
	newResults := []int{}

	for _, node := range nodes {
		if node == nil {
			continue
		}
		newResults = append(newResults, node.Val)
		newNodes = append(newNodes, node.Left)
		newNodes = append(newNodes, node.Right)
	}

	if len(newResults) != 0 {
		result.Data = append(result.Data, newResults)
	}
	levelOrderHelper(newNodes, result)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := &levelOrderData{Data: [][]int{}}

	levelOrderHelper([]*TreeNode{root}, result)
	return result.Data
}
