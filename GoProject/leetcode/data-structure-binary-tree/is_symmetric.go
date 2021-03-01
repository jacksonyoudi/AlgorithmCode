package main

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
