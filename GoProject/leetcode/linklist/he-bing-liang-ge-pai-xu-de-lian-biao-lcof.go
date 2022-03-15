package linklist

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var first *ListNode = new(ListNode)

	head := first

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			first.Next = l1
			first = l1
			l1 = l1.Next
		} else {
			first.Next = l2
			first = l2
			l2 = l2.Next
		}
	}
	for l1 != nil {
		first.Next = l1
		first = l1
		l1 = l1.Next
	}

	for l2 != nil {
		first.Next = l2
		first = l2
		l2 = l2.Next
	}

	return head.Next

}
