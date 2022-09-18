package merge_two_sorted_lists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// dummy
	dummy := &ListNode{-1, nil}
	currey := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			currey.Next = p1
			p1 = p1.Next
		} else {
			currey.Next = p2
			p2 = p2.Next
		}

		currey = currey.Next
		//可以去掉
		currey.Next = nil
	}

	if p1 != nil {
		currey.Next = p1
	}

	if p2 != nil {
		currey.Next = p2
	}

	return dummy.Next

}
