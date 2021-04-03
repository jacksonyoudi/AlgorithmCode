package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	fast, slow := head.Next.Next, head.Next

	for fast != slow {
		//  如果没有环
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast

}
