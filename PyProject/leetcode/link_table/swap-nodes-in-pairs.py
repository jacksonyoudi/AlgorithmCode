# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        thead = ListNode(-1)
        thead.next = head
        cur = thead
        while cur.next and cur.next.next:
            # swap
            a, b = cur.next, cur.next.next
            cur.next, a.next = b, b.next
            b.next = a
            cur = cur.next.next

        return thead.next
