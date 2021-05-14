package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nodes []*TreeNode, result [][]int) {
	if len(nodes) == 0 {
		return
	}

	newNodes := make([]*TreeNode, 0)
	newResults := make([]int, 0)
	for _, node := range nodes {
		if node != nil {
			newResults = append(newResults, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}

		result = append([][]int{newResults}, result...)

		helper(newNodes, result)
	}

}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	helper([]*TreeNode{root}, result)
	return result
}
