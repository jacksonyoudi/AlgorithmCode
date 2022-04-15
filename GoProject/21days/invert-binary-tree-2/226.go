package invert_binary_tree_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	right := invertTree(root.Left)
	left := invertTree(root.Right)

	root.Left = left
	root.Right = right
	return root
}
