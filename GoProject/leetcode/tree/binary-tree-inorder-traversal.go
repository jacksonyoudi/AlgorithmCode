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

//func inorderTraversalOfHelper(root *TreeNode, res []int) []int {
//	if root == nil {
//		return res
//	}
//	res = inorderTraversalOfHelper(root.Left, res)
//	res = append(res, root.Val)
//	res = inorderTraversalOfHelper(root.Right, res)
//	return res
//}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
