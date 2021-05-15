package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nodes []*TreeNode, data *Data) {
	if len(nodes) == 0 {
		return
	}

	newNodes := make([]*TreeNode, 0)
	newResults := make([]int, 0)
	for _, node := range nodes {
		if node != nil {
			newResults = append(newResults, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
	}
	if len(newResults) > 0 {
		data.D = append([][]int{newResults}, data.D...)
	}

	if len(newNodes) > 0 {
		helper(newNodes, data)
	}
}

type Data struct {
	D [][]int
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)

	data := &Data{
		D: result,
	}

	helper([]*TreeNode{root}, data)
	return data.D
}

func main() {
	root := &TreeNode{1, nil, nil}
	a := levelOrderBottom(root)
	fmt.Println(a)

}
