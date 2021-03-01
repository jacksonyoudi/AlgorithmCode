package main

func pathSum(root *TreeNode, targetSum int, path int) bool {
	if root == nil {
		return false
	}

	p := path + root.Val

	if root.Right == nil && root.Left == nil {
		if targetSum == p {
			return true
		} else {
			return false
		}
	}

	return pathSum(root.Left, targetSum, p) || pathSum(root.Right, targetSum, p)

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(root, targetSum, 0)
}
