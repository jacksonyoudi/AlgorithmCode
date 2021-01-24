package tree

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

func inorderTraversalOfHelper(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inorderTraversalOfHelper(root.Left, res)
	res = append(res, root.Val)
	res = inorderTraversalOfHelper(root.Right, res)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalOfHelper(root, []int{})
}
