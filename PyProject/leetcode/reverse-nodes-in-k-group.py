from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class ReverseNodesInKGroup:
    def reverse(self, start: ListNode, end: ListNode) -> ListNode:
        pre = None
        cur = start

        while cur != end:
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        return pre

    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if head is None:
            return head
        a = b = head
        for i in range(k):
            if b is None:
                return head
            b = b.next

        new = self.reverse(a, b)
        a.next = self.reverseKGroup(b, k)
        return new
