package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	result := reversePrint(head.Next)

	result = append(result, head.Val)

	return result

}
