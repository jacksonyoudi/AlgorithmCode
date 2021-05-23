package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(nums []*ListNode, start, end int) int {
	pivot := nums[end].Val
	i := start - 1
	for j := start; j < end; j++ {
		if nums[j].Val < pivot {
			i += 1
			nums[i].Val, nums[j].Val = nums[j].Val, nums[i].Val
		}
	}
	nums[i+1].Val, nums[end].Val = nums[end].Val, nums[i+1].Val
	return i + 1
}

func quick_sort_list(nums []*ListNode, start, end int) {
	if start < end {
		par := partition(nums, start, end)
		quick_sort_list(nums, start, par-1)
		quick_sort_list(nums, par+1, end)
	}

}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	lists := make([]*ListNode, 0)

	for head != nil {
		lists = append(lists, head)
		head = head.Next
	}

	quick_sort_list(lists, 0, len(lists)-1)

	return lists[0]

}
