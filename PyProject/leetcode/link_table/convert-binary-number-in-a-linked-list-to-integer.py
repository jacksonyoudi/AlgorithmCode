# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getDecimalValue(self, head: ListNode) -> int:
        if head is None:
            return 0
        result = []
        while head:
            result.append(head.val)
            head = head.next
        r = 0
        for i, v in enumerate(result[::-1]):
            r += v * pow(2, i)

        return r
