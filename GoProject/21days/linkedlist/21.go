package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsHelper(header, list1, list2 *ListNode) {
	if list1 == nil {
		header.Next = list2
		return
	}

	if list2 == nil {
		header.Next = list1
		return
	}

	if list1.Val < list2.Val {
		header.Next = list1
		next := list1.Next
		header.Next.Next = nil
		mergeTwoListsHelper(header.Next, next, list2)
	} else {
		header.Next = list2
		next := list2.Next
		header.Next.Next = nil
		mergeTwoListsHelper(header.Next, list1, next)
	}

}

// 使用递归方式
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}

	mergeTwoListsHelper(dummy, list1, list2)
	return dummy.Next
}
