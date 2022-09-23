package flatten_binary_tree_to_linked_list

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

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := root.Left
	right := root.Right

	root.Right = left

}

func flatten(root *TreeNode) {

}
