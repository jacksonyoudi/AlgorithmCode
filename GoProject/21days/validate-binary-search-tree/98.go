package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTHelper(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return isValidBSTHelper(root.Left, min, root) && isValidBSTHelper(root.Right, root, max)

}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}
