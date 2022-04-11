package remove_nth_node_from_end_of_list_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head

	for n > 0 {
		if right == nil {
			return head
		}
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return head
}
