package linkedlist

/**
 * Definition for singly-linked list.
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// merge

	nextVal := 0
	head := &ListNode{
		-1,
		nil,
	}

	current := head

	l1Next := l1
	l2Next := l2

	for l1Next != nil || l2Next != nil {
		one := 0
		two := 0

		if l1Next != nil {
			one = l1Next.Val
			l1Next = l1Next.Next
		}

		if l2Next != nil {
			two = l2Next.Val
			l2Next = l2Next.Next
		}

		currentVal := one + two + nextVal
		nextVal = 0
		if currentVal >= 10 {
			nextVal = 1
			currentVal = currentVal - 10
		}

		current.Next = &ListNode{Val: currentVal, Next: nil}
		current = current.Next

	}

	if nextVal == 1 {
		current.Next = &ListNode{Val: 1, Next: nil}
	}

	return head.Next
}
