package tree

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func helper(root *Node, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if len(root.Children) == 0 {
		return res
	}
	for _, node := range root.Children {
		res = helper(node, res)
	}
	return res
}

func preorder(root *Node) []int {
	return helper(root, []int{})
}
