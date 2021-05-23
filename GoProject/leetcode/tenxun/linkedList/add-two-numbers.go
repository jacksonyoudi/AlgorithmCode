package linkedList

/*type ListNode struct {
	Val  int
	Next *ListNode
} */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: -1, Next: nil}
	head := root

	addOne := 0
	for l1 != nil && l2 != nil {

		i := l1.Val + l2.Val + addOne
		addOne = 0

		if i >= 10 {
			i = i - 10
			addOne = 1
		}
		node := &ListNode{Val: i, Next: nil}

		head.Next = node
		head = node
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {

		i := l1.Val + addOne
		addOne = 0
		if i >= 10 {
			i = i - 10
			addOne = 1
		}
		node := &ListNode{Val: i, Next: nil}

		head.Next = node
		head = node
		l1 = l1.Next
	}

	for l2 != nil {

		i := l2.Val + addOne
		addOne = 0
		if i >= 10 {
			i = i - 10
			addOne = 1
		}
		node := &ListNode{Val: i, Next: nil}

		head.Next = node
		head = node
		l2 = l2.Next
	}

	if addOne == 1 {
		head.Next = &ListNode{Val: 1, Next: nil}
		addOne = 0
	}

	return root.Next

}
