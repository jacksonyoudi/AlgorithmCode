package tree

func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return
}
