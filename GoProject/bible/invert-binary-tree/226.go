package invert_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	root.Right, root.Left = root.Left, root.Right
	helper(root.Left)
	helper(root.Right)

}

func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}
