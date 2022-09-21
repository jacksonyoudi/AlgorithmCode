package flatten_binary_tree_to_linked_list

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func helper(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	helper(node1.Left, node1.Right)
	helper(node2.Left, node2.Right)

	helper(node1.Right, node2.Left)
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	helper(root.Left, root.Right)
	return root

}
