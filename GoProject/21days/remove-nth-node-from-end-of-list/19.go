package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{
		Val:  -1,
		Next: head,
	}
	fast := dump
	slow := dump

	for i := 0; i < n; i++ {
		if fast == nil {
			return dump.Next
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head.Next
}
