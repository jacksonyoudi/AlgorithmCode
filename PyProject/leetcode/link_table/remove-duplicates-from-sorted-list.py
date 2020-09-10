# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head:
            return head
        pre = head
        after = head.next
        while after:
            if pre.val == after.val:
                after = after.next
                pre.next = after
            else:
                pre = pre.next
                after = after.next

        return head
