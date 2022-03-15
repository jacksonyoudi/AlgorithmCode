package main

import "fmt"

// 通过数据结构进行构造成树，然后再配合bfs

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

type Tree struct {
	nodes    map[int]*NTreeNode
	roots    []*NTreeNode
	nodeVals map[int]int
}

func (t *Tree) Add(n []int) {
	one := n[0]
	two := n[1]
	oneNode, ok := t.nodes[one]
	if !ok {
		oneNode = &NTreeNode{
			Val:      one,
			Children: make([]*NTreeNode, 0),
		}
		t.nodes[one] = oneNode
		t.nodeVals[one] = 1
	}
	t.nodes[one] = oneNode

	twoNode, ok := t.nodes[two]
	if !ok {
		twoNode = &NTreeNode{
			Val:      two,
			Children: make([]*NTreeNode, 0),
		}
		t.nodes[two] = twoNode
		t.nodeVals[two] = 1
	}

	twoNode.Children = append(twoNode.Children, oneNode)
	t.nodes[two] = twoNode
}

func (t *Tree) FlushRoot() {

	allVals := t.nodeVals

	// 找到开头的课程
	// 如果有子节点的， 子节点就跳过
	// 先清空
	t.roots = make([]*NTreeNode, 0)

	for _, node := range t.nodes {
		if len(node.Children) != 0 {
			for _, n := range node.Children {
				delete(allVals, n.Val)
			}
		}
	}
	for j, _ := range allVals {
		tmp := t.nodes[j]
		t.roots = append(t.roots, tmp)
	}
}

func (t *Tree) Roots() []*NTreeNode {
	return t.roots
}

func helperdfs(root *NTreeNode, depth int, result []int) ([]int, bool) {
	if depth == 0 {
		return result, true
	}

	if root == nil {
		return []int{}, false
	}

	newResult := append(result, root.Val)

	if len(root.Children) == 0 {
		if depth == 1 {
			return newResult, true
		} else {
			return []int{}, false
		}
	} else {
		for _, child := range root.Children {
			r, ok := helperdfs(child, depth-1, newResult)
			if ok {
				return r, ok
			}
		}
	}

	return []int{}, false

}

func (t *Tree) bfs(depth int) []int {
	result := make([]int, 0)
	for _, root := range t.roots {
		r, ok := helperdfs(root, depth, result)
		if ok {
			return r
		}
	}
	return []int{}
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 边界条件
	if numCourses <= 0 {
		return []int{}
	}

	// 构造树
	tree := &Tree{
		nodes:    make(map[int]*NTreeNode),
		roots:    make([]*NTreeNode, 0),
		nodeVals: make(map[int]int),
	}

	for _, prerequisite := range prerequisites {
		tree.Add(prerequisite)
	}
	tree.FlushRoot()

	return tree.bfs(numCourses)

}

func main() {

	a := []int{1, 0}
	ints := findOrder(2, [][]int{a})
	fmt.Println(ints)
}
