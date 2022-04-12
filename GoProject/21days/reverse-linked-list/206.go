package reverse_linked_list

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

func reverseList(head *ListNode) *ListNode {
	_, root := helper(head)
	return root
}
