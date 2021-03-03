package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := maxDepth(root.Right) + 1
	left := maxDepth(root.Left) + 1
	if right > left {
		return right
	}
	return left

}
