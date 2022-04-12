package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func helper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	node, root := helper(head.Next)
	node.Next = head
	head.Next = nil
	return head, root

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dump := &ListNode{-1, head}

	i := 0
	start, end := dump, dump
	for i < right {
		if i < left-1 {
			if start == nil || start.Next == nil {
				return head
			} else {
				start = start.Next
			}
		}

		if end == nil || end.Next == nil {
			return head
		}
		end = end.Next
		i++
	}
	cur := end.Next
	end.Next = nil

	node, root := helper(start.Next)
	start.Next = root
	node.Next = cur
	return dump.Next
}
