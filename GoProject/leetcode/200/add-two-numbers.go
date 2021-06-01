package _200

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{Val: -1, Next: nil}

	tmp := head

	one, two := l1, l2
	flag := 0

	for one != nil && two != nil {
		v := one.Val + two.Val + flag
		flag = 0
		if v >= 10 {
			v = v - 10
			flag = 1
		}

		one = one.Next
		two = two.Next

		node := &ListNode{Val: v, Next: nil}
		tmp.Next = node
		tmp = node
	}

	for one != nil {
		v := one.Val + flag
		flag = 0
		if v >= 10 {
			v = v - 10
			flag = 1
		}
		one = one.Next

		node := &ListNode{Val: v, Next: nil}
		tmp.Next = node
		tmp = node

	}

	for two != nil {
		v := two.Val + flag
		flag = 0
		if v >= 10 {
			v = v - 10
			flag = 1
		}
		two = two.Next

		node := &ListNode{Val: v, Next: nil}
		tmp.Next = node
		tmp = node
	}

	if flag == 1 {
		node := &ListNode{Val: 1, Next: nil}
		tmp.Next = node
		tmp = node
	}

	return head.Next

}
