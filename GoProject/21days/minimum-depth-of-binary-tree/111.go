package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nodes []*TreeNode, depth int) int {
	if len(nodes) == 0 {
		return depth
	}

	cur := []*TreeNode{}

	for _, node := range nodes {
		if node == nil {
			return depth
		}
		if node.Right == nil && node.Left == nil {
			return depth
		}

		if node.Left != nil {
			cur = append(cur, node.Left)
		}
		if node.Right != nil {
			cur = append(cur, node.Right)
		}
	}

	if len(cur) == 0 {
		return depth
	}

	return helper(cur, depth+1)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper([]*TreeNode{root}, 1)
}
