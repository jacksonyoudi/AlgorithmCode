package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectHelper(one, two *Node) {
	if one == nil || two == nil {
		return
	}
	// 当前
	one.Next = two

	// one, two作为父节点
	connectHelper(one.Left, one.Right)
	connectHelper(two.Left, two.Right)

	// one two共同父节点
	connectHelper(one.Right, two.Left)

}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectHelper(root.Left, root.Right)
	return root
}
