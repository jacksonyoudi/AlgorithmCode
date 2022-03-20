package delete_node_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left != nil && root.Right != nil {

			return root.Left
		} else {

		}
	}

	return root

}
