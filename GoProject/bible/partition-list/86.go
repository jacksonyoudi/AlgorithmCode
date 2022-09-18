package partition_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 小于
	l1 := &ListNode{-1, nil}
	p1 := l1

	// 大于
	l2 := &ListNode{-1, nil}
	p2 := l2

	currency := head

	for currency != nil {
		if currency.Val < x {
			p1.Next = currency
			p1 = p1.Next
			currency = currency.Next
			p1.Next = nil
		} else {
			p2.Next = currency
			p2 = p2.Next
			currency = currency.Next
			p2.Next = nil
		}
	}

	if l1.Next != nil {
		p1.Next = l2.Next
	} else if l2.Next != nil {
		return l2.Next
	} else {
		return nil
	}

	return l1.Next
}
