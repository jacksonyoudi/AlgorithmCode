package dfs

type KthSmallestData struct {
	data []int
}

func kthSmallestDfs(root *TreeNode, d *KthSmallestData, k int) {
	if root == nil {
		return
	}

	kthSmallestDfs(root.Left, d, k)
	d.data = append(d.data, root.Val)
	if len(d.data) == k {
		return
	}

	kthSmallestDfs(root.Right, d, k)

}

// 中序遍历
func kthSmallest(root *TreeNode, k int) int {

	d := &KthSmallestData{
		data: make([]int, 0),
	}

	kthSmallestDfs(root, d, k)

	return d.data[k-1]
}
