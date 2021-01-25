package tree


// 自底向上
func maxHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxHelper(root.Left)
	right := maxHelper(root.Right)
	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

func maxDepth(root *TreeNode) int {
	return maxHelper(root)
}
