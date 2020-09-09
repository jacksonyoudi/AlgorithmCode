# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        if not head:
            return None
        slow = head
        quick = head
        while quick:
            if k >= -1:
                k -= 1
                quick = quick.next
            else:
                quick = quick.next
                slow = slow.next
        return slow
