package bfs

//type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
//}

type D struct {
	Data []int
}

func get(roots []*TreeNode, result *D) {
	if len(roots) == 0 {
		return
	}
	nodes := make([]*TreeNode, 0)

	a := result.Data

	for _, root := range roots {
		if root != nil {
			a = append(a, root.Val)
			if root.Left != nil {
				nodes = append(nodes, root.Left)
			}
			if root.Right != nil {
				nodes = append(nodes, root.Right)
			}
		}
	}

	result.Data = a
	get(nodes, result)

}

func levelOrder(root *TreeNode) []int {
	n := make([]int, 0)
	d := &D{
		Data: n,
	}

	get([]*TreeNode{root}, d)
	return d.Data
}
