# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        if not head:
            return head
        quick = head
        slow = head

        while quick:
            if quick.next:
                quick = quick.next.next
                slow = slow.next
        
