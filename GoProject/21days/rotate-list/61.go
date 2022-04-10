package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	n := 0
	dump := &ListNode{
		Val:  -1,
		Next: head,
	}
	tmp := head
	end := head
	for tmp != nil {
		n++
		if tmp.Next == nil {
			end = tmp
		}
		tmp = tmp.Next
	}

	k = n - k%n

	if k == 0 {
		return head
	}

	tp := head
	for i := 0; i < k; i++ {
		tp = tp.Next
	}
	dump.Next = tp.Next
	tp.Next = nil

	end.Next = head

	return dump.Next

}
