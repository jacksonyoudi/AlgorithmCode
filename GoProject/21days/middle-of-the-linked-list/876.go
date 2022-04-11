package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}
	return left
}
