package linklist

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	one := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return one
}
