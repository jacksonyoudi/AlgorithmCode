from typing import List


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def helper(l1: ListNode, l2: ListNode, header: ListNode):
    if l1 is None:
        header.next = l2
        return
    if l2 is None:
        header.next = l1
        return
    if l1.val < l2.val:
        header.next = ListNode(l1.val)
        header = header.next
        helper(l1.next, l2, header)
    else:
        header.next = ListNode(l2.val)
        header = header.next
        helper(l1, l2.next, header)


class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        if l2 is None:
            return l1

        if l1.val < l2.val:
            header = ListNode(l1.val)
            helper(l1.next, l2, header)
        else:
            header = ListNode(l2.val)
            helper(l1, l2.next, header)
        return header
