package main

func level(nodes []*TreeNode, result [][]int) [][]int {

	if len(nodes) == 0 {
		return result
	}

	res := []int{}
	newNodes := make([]*TreeNode, 0)
	for _, node := range nodes {
		if node != nil {
			res = append(res, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
	}

	if len(res) == 0 {
		return result
	}
	result = append(result, res)
	return level(newNodes, result)

}

func levelOrder(root *TreeNode) [][]int {
	return level([]*TreeNode{root}, make([][]int, 0))
}
