package dfs

func minDepth(root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return 1
	}

	right := minDepth(root.Right) + 1
	left := minDepth(root.Left) + 1
	if right > left {
		return left
	}
	return right

}
