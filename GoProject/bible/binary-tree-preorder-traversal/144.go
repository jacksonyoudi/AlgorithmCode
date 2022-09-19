package binary_tree_preorder_traversal

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

type T struct {
	Data []int
}

func helper(root *TreeNode, t *T) {
	if root == nil {
		return
	}
	t.Data = append(t.Data, root.Val)
	helper(root.Left, t)
	helper(root.Right, t)

}
func preorderTraversal(root *TreeNode) []int {
	t := &T{Data: make([]int, 0)}
	helper(root, t)
	return t.Data

}
