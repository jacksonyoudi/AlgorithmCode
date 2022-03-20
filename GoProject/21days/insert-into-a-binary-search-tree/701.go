package insert_into_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBSTHelper(root *TreeNode, val int) {
	// 都是使用left
	if root.Val >= val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &TreeNode{val, nil, nil}
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &TreeNode{val, nil, nil}
		}
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//if root == nil {
	//	return &TreeNode{val, nil, nil}
	//}
	//insertIntoBSTHelper(root, val)
	//return root

	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	// 不会有重复的数据
	//if root.Val == val {
	//}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}
