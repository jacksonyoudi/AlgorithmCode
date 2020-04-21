class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def reverse_list(head: ListNode) -> ListNode:
    if head is None or head.next is None:
        return head
    new_list = reverse_list(head.next)
    t1 = head.next
    t1.next = head
    head.next = None
    return new_list
