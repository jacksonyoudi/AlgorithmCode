# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class LinkedListCycleII:
    def detectCycle(self, head: ListNode) -> ListNode:
        s = set()
        while head is not None:
            if head in s:
                return head
            else:
                s.add(head)
            head = head.next
        return None
