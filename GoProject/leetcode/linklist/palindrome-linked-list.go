package linklist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}
	slow, quick := head, head.Next

	for slow != nil && quick != nil {
		if slow == quick {
			return true
		}
		slow = slow.Next
		if quick.Next != nil {
			quick = quick.Next.Next
		} else {
			return false
		}
	}

	return false
}
