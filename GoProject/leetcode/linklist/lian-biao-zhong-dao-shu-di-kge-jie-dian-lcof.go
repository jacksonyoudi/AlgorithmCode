package linklist

func getKthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	l := 0
	for head.Next != nil {
		l += 1
		head = head.Next
	}

	head = first
	if l < n {
		return head
	}

	m := l - n
	for j := 0; j < m; j++ {
		head = head.Next
	}
	head.Val = head.Next.Val
	head.Next = head.Next.Next

	return head

}
