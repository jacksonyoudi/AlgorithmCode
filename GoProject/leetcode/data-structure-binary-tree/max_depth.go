package main

func max_depth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	right, left := depth+1, depth+1

	if root.Right != nil {
		right = max_depth(root.Right, depth+1)
	}

	if root.Left != nil {
		left = max_depth(root.Left, depth+1)
	}

	if right > left {
		return right
	} else {
		return left
	}

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	a, b := 1, 1
	if root.Right != nil {
		a = maxDepth(root.Right) + 1
	}

	if root.Left != nil {
		b = maxDepth(root.Left) + 1
	}

	if a > b {
		return a
	} else {
		return b
	}
}
