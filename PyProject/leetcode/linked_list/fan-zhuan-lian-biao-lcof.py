# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if not head and not head.next:
            return head
        # after = head.next
        # cur = self.reverseList(after)
        # head.next.next = head
        # head.next = None
        # return cur

        cur = head

        while cur and cur.next:
            prev = cur
            nex = cur.next.next
            prev.next = nex
            cur.next = prev
            cur = nex
        return cur
