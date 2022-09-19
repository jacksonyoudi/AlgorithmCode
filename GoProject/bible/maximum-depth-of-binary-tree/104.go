package maximum_depth_of_binary_tree

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		a := maxDepth(root.Right)
		b := maxDepth(root.Left)
		if a > b {
			return a + 1
		} else {
			return b + 1
		}
	}
}
