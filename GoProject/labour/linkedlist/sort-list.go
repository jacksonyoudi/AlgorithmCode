package linkedlist

func partition(nums []int, start int, end int) int {
	pivot := nums[end]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j] <= pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[end], nums[i+1] = nums[i+1], nums[end]
	return i + 1

}

func quick_sort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quick_sort(nums, start, pivot-1)
		quick_sort(nums, pivot+1, end)
	}

}

func sortList(head *ListNode) *ListNode {

	// 边界
	if head == nil || head.Next == nil {
		return head
	}

	tmp := make([]int, 0)
	l1Next := head
	for l1Next != nil {
		tmp = append(tmp, l1Next.Val)
		l1Next = l1Next.Next
	}

	quick_sort(tmp, 0, len(tmp)-1)

	current := head
	for _,v := range tmp {
		current.Val = v
		current = current.Next
	}
	return head
}
