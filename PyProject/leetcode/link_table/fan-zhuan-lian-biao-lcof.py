# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


#     ListNode reverse = reverseList(head.next);
#     head.next.next = head;
#     head.next = null;
#     return reverse;
#

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if not head or head.next:
            return head
        after = head.next
        cur = self.reverseList(after)
        cur.next = head
        head.next = None
        return cur



