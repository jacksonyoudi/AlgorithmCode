package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: -1, Next: nil}
	head := root

	for l1 != nil && l2 != nil {

		v := 0
		if l1.Val < l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else {
			v = l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{Val: v, Next: nil}
		head = head.Next
	}

	for l1 != nil {

		v := l1.Val
		l1 = l1.Next

		head.Next = &ListNode{Val: v, Next: nil}
		head = head.Next
	}

	for l2 != nil {

		v := l2.Val
		l2 = l2.Next
		head.Next = &ListNode{Val: v, Next: nil}
		head = head.Next
	}

	return root.Next

}
