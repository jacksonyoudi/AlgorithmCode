package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func helperReverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	node, root := helperReverseList(head.Next)
	node.Next = head
	head.Next = nil
	return head, root

}

func reverseList(head *ListNode) *ListNode {
	_, r := helperReverseList(head)
	return r
}
