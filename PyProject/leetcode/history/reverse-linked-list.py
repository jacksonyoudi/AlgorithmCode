# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class ReverseLinkedList:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None:
            return head
        pre_node = None
        while head:
            next_node = head.next
            head.next = pre_node
            pre_node = head
            head = next_node
        return pre_node
