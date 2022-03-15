package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
