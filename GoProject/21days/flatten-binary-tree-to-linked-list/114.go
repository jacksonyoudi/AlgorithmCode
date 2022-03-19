package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		return
	}

	right := root.Right
	left := root.Left
	flatten(root.Right)
	flatten(root.Left)

	root.Left = nil
	root.Right = left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
